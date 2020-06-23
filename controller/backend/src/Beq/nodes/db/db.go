package db

import (
	"Beq/nodes/model"
	ruleDB "Beq/rules/db"
	setting "Beq/settings/db"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type db map[string]model.NodeData

var instance db
var allNodes map[string]int
var nodeNameMap map[string]string
var ipMacMap map[string]string
var nodeIndex int
var connectivityChecker map[string]bool
var flowChecker map[string]bool
var once sync.Once

//GetDataBase Initiating list database
func GetDataBase() *db {
	once.Do(func() {
		instance = make(map[string]model.NodeData)
		allNodes = make(map[string]int)
		ipMacMap = make(map[string]string)
		nodeNameMap = make(map[string]string)
		connectivityChecker = make(map[string]bool)
		flowChecker = make(map[string]bool)
		nodeIndex = 0
	})
	return &instance
}

func (*db) AddNode(MAC string, NodeData model.NodeData) {
	if instance != nil {
		nodeNameMap[MAC] = NodeData.Node.Name
		ipMacMap[NodeData.Node.IP] = NodeData.Node.MAC
		instance[MAC] = NodeData
		if allNodes[MAC] == 0 {
			allNodes[MAC] = nodeIndex
			nodeIndex++
		}
		for _, NodeMAC := range NodeData.Neighbours {
			allNodes[NodeMAC.MAC] = nodeIndex
			nodeIndex++
		}
	}
}

func (*db) GetMacForIP(IP string) string {
	return ipMacMap[IP]
}

func (*db) GenarateNetworkTopology() (map[string]string, model.GrpData, error) {
	connectivityChecker = make(map[string]bool)
	GrpData := model.GrpData{}
	if instance == nil {
		return nil, GrpData, errors.New("No Node Data")
	}

	nodes := []model.GrpNode{}
	config := setting.GetSystemSetting()
	controllerMAC, err := config.GetMAC()
	if err != nil {
		return nil, GrpData, errors.New("No Node Data")
	}

	for MAC, Index := range allNodes {
		a := instance[MAC]
		group := "AP"

		if a.Node.IP == "" {
			if MAC == controllerMAC {
				group = "Controller"
			} else {
				group = "NotAP"
			}
		}

		nodes = append(nodes,
			model.GrpNode{
				ID:       Index,
				Label:    a.Node.Name,
				Group:    group,
				NodeData: instance[MAC],
			})

	}

	nodeLinks := []model.GrpNodeLink{}
	for MAC, NadeData := range instance {
		for _, Neighbour := range NadeData.Neighbours {

			if isConnected(allNodes[MAC], allNodes[Neighbour.MAC]) == false {
				makeConectivity(allNodes[MAC], allNodes[Neighbour.MAC])
				currentLink := model.GrpNodeLink{}
				isCotrollerLink := false
				if MAC == controllerMAC || Neighbour.MAC == controllerMAC {
					isCotrollerLink = true
				}
				currentLink.SetLink(allNodes[MAC], allNodes[Neighbour.MAC], Neighbour.Bandwidth, isCotrollerLink)
				nodeLinks = append(nodeLinks, currentLink)
			}

		}

	}
	GrpData.Nodes = nodes
	GrpData.Edges = nodeLinks
	return nodeNameMap, GrpData, nil
}

func (*db) GenarateNetworkTopologyWithFlowHighlight() (map[string]string, model.GrpData, error) {
	connectivityChecker = make(map[string]bool)
	GrpData := model.GrpData{}
	if instance == nil {
		return nil, GrpData, errors.New("No Node Data")
	}
	//get Flow links and put in flow links array
	flowLinks, err := getNodeLinksByFlowID("F0")
	if err != nil {
		return nil, GrpData, errors.New("Error in Flow generation process")
	}
	for _, s := range flowLinks {

		makeFlow(s.DstNode, s.SrcNode)
	}

	nodes := []model.GrpNode{}
	config := setting.GetSystemSetting()
	controllerMAC, err := config.GetMAC()
	if err != nil {
		return nil, GrpData, errors.New("No Node Data")
	}

	for MAC, Index := range allNodes {
		a := instance[MAC]
		group := "AP"

		if a.Node.IP == "" {
			if MAC == controllerMAC {
				group = "Controller"
			} else {
				group = "NotAP"
			}
		}

		nodes = append(nodes,
			model.GrpNode{
				ID:       Index,
				Label:    a.Node.Name,
				Group:    group,
				NodeData: instance[MAC],
			})

	}

	nodeLinks := []model.GrpNodeLink{}
	for MAC, NadeData := range instance {
		for _, Neighbour := range NadeData.Neighbours {

			if isConnected(allNodes[MAC], allNodes[Neighbour.MAC]) == false {
				makeConectivity(allNodes[MAC], allNodes[Neighbour.MAC])
				edgeColor, isDashed := isFlowEdge(allNodes[MAC], allNodes[Neighbour.MAC])
				fmt.Println(edgeColor, isDashed)
				currentLink := model.GrpNodeLink{}
				isCotrollerLink := false
				if MAC == controllerMAC || Neighbour.MAC == controllerMAC {
					isCotrollerLink = true
				}
				currentLink.SetLinkWithColor(
					allNodes[MAC],
					allNodes[Neighbour.MAC],
					Neighbour.Bandwidth,
					isCotrollerLink,
					edgeColor,
					isDashed,
				)
				nodeLinks = append(nodeLinks, currentLink)
			}

		}

	}
	GrpData.Nodes = nodes
	GrpData.Edges = nodeLinks
	return nodeNameMap, GrpData, nil
}

//GetNodeLinksByFlowID used for remove Rule by RuleID
func getNodeLinksByFlowID(FlowID string) ([]model.FlowLink, error) {
	flowLinks := []model.FlowLink{}
	for _, RuleData := range *ruleDB.GetRuleStore() {
		if RuleData.FlowID == FlowID && RuleData.Action == "ACCEPT" {
			flowLinks = append(flowLinks, model.FlowLink{
				DstNode: allNodes[RuleData.DstMAC], SrcNode: allNodes[instance.GetMacForIP(RuleData.NodeIP)],
			})
		}
	}
	fmt.Println(flowLinks)
	return flowLinks, nil

}

func makeConectivity(node1Id int, node2Id int) {
	if node1Id != node2Id {
		maxID := 0
		minID := 0
		if node1Id > node2Id {
			maxID = node1Id
			minID = node2Id
		} else {
			maxID = node1Id
			minID = node2Id
		}
		KEY := strconv.Itoa(maxID) + "edge" + strconv.Itoa(minID)
		connectivityChecker[KEY] = true
	}
}

func isConnected(node1Id int, node2Id int) bool {
	if node1Id != node2Id {
		maxID := 0
		minID := 0
		if node1Id > node2Id {
			maxID = node1Id
			minID = node2Id
		} else {
			minID = node1Id
			maxID = node2Id
		}
		KEY := strconv.Itoa(maxID) + "edge" + strconv.Itoa(minID)
		if connectivityChecker[KEY] {
			return true
		}
		return false

	}
	return true
}

func makeFlow(node1Id int, node2Id int) {
	if node1Id != node2Id {
		maxID := 0
		minID := 0
		if node1Id > node2Id {
			maxID = node1Id
			minID = node2Id
		} else {
			maxID = node1Id
			minID = node2Id
		}
		KEY := strconv.Itoa(maxID) + "flowEdge" + strconv.Itoa(minID)
		fmt.Println("--------------")
		flowChecker[KEY] = true
		fmt.Println(flowChecker)
		fmt.Println("--------------")
	}
}

//check whather this edge is belongs to flow
func isFlowEdge(node1Id int, node2Id int) (string, bool) {
	if node1Id != node2Id {
		maxID := 0
		minID := 0
		if node1Id > node2Id {
			maxID = node1Id
			minID = node2Id
		} else {
			minID = node1Id
			maxID = node2Id
		}
		KEY := strconv.Itoa(maxID) + "flowEdge" + strconv.Itoa(minID)
		fmt.Println(KEY, flowChecker[KEY])
		if flowChecker[KEY] == true {
			return "green", false
		}
		return "gray", true

	}
	return "gray", true
}
