package db

import (
	"Beq/nodes/model"
	"errors"
	"strconv"
	"sync"
)

type db map[string]model.NodeData

var instance db
var allNodes map[string]int
var nodeIndex int

var once sync.Once

//GetDataBase Initiating list database
func GetDataBase() *db {
	once.Do(func() {
		instance = make(map[string]model.NodeData)
		allNodes = make(map[string]int)
		nodeIndex = 0
	})
	return &instance
}

func (*db) AddNode(MAC string, NodeData model.NodeData) {
	if instance != nil {
		instance[MAC] = NodeData
		allNodes[MAC] = nodeIndex
		nodeIndex++
		for _, NodeMAC := range NodeData.Neighbours {
			allNodes[NodeMAC.MAC] = nodeIndex
			nodeIndex++
		}
	}
}

func (*db) GenarateNetworkTopology() (model.GrpData, error) {
	GrpData := model.GrpData{}
	if instance == nil {
		return GrpData, errors.New("No Node Data")
	}

	nodes := []model.GrpNode{}
	nodeIndex := 0

	for _, Index := range allNodes {
		nodes = append(nodes, model.GrpNode{ID: Index, Label: "Node " + strconv.Itoa(Index), Group: "Agent"})
		nodeIndex++
	}

	nodeLinks := []model.GrpNodeLink{}
	for MAC, NadeData := range instance {
		for _, Neighbour := range NadeData.Neighbours {
			currentLink := model.GrpNodeLink{}
			currentLink.SetLink(allNodes[MAC], allNodes[Neighbour.MAC],strconv.Itoa(Neighbour.Bandwidth) + " mbps")
			nodeLinks = append(nodeLinks, currentLink)
		}

	}
	GrpData.Nodes = nodes
	GrpData.Edges = nodeLinks
	return GrpData, nil
}
