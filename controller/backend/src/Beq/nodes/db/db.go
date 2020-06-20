package db

import (
	"Beq/nodes/model"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type db map[string]model.NodeData

var instance db
var allNodes map[string]int
var nodeNameMap map[string]string
var nodeIndex int
var connectivityChecker map[string]bool
var once sync.Once

//GetDataBase Initiating list database
func GetDataBase() *db {
	once.Do(func() {
		instance = make(map[string]model.NodeData)
		allNodes = make(map[string]int)
		nodeNameMap = make(map[string]string)
		connectivityChecker = make(map[string]bool)
		nodeIndex = 0
	})
	return &instance
}

func (*db) AddNode(MAC string, NodeData model.NodeData) {
	if instance != nil {
		nodeNameMap[MAC] = NodeData.Node.Name
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

func (*db) GenarateNetworkTopology() (model.GrpData, error) {
	connectivityChecker = make(map[string]bool)
	GrpData := model.GrpData{}
	if instance == nil {
		return GrpData, errors.New("No Node Data")
	}

	nodes := []model.GrpNode{}

	for MAC, Index := range allNodes {
		a := instance[MAC]
		group := "AP"
		if a.Node.IP == "" {
			group = "NotAP"
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
				currentLink.SetLink(allNodes[MAC], allNodes[Neighbour.MAC], Neighbour.Bandwidth)
				nodeLinks = append(nodeLinks, currentLink)
			}

		}

	}
	GrpData.Nodes = nodes
	GrpData.Edges = nodeLinks
	return GrpData, nil
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
		fmt.Println(KEY)
		if connectivityChecker[KEY] {
			return true
		}
		return false

	}
	return true
}
