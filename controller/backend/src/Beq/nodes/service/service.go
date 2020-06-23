package service

import (
	"Beq/nodes/db"
	"Beq/nodes/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AddNodeInfo use update node informations
func AddNodeInfo(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	neighbourMap := db.GetDataBase()
	var NodeData model.NodeData
	reqBody, err := ioutil.ReadAll(r.Body)

	resp := model.Response{}
	// // set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := json.Unmarshal(reqBody, &NodeData)
		if err != nil {
			log.Println("ERROR: Payload Error", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			neighbourMap.AddNode(NodeData.Node.MAC, NodeData)
			log.Println("INFO: [NO]: Node :", NodeData.Node.Name, " info added. IP : ", NodeData.Node.IP)
			resp.Code = http.StatusOK
			resp.Message = "Data Base Updated"
			resp.Data = nil
			w.WriteHeader(http.StatusOK)
		}
	}
	json.NewEncoder(w).Encode(resp)

}

// GetNodeInfo for genarate information required for Network Topology graph
func GetNodeInfo(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	neighbourMap := db.GetDataBase()
	resp := model.Response{}
	// set Default value
	resp.Default()

	NodeNameMap, GraphData, err := neighbourMap.GenarateNetworkTopology()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data := make(map[string]interface{})
		data["graphData"] = GraphData
		data["nodeNames"] = NodeNameMap

		resp.Code = http.StatusOK
		resp.Message = "Updated time: "
		resp.Data = data
		w.WriteHeader(http.StatusOK)

	}
	json.NewEncoder(w).Encode(resp)

}

// GetNodeInfoWithFlowHighlight for genarate information required for Network Topology graph
func GetNodeInfoWithFlowHighlight(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	neighbourMap := db.GetDataBase()
	resp := model.Response{}
	// set Default value
	resp.Default()

	NodeNameMap, GraphData, err := neighbourMap.GenarateNetworkTopologyWithFlowHighlight()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data := make(map[string]interface{})
		data["graphData"] = GraphData
		data["nodeNames"] = NodeNameMap

		resp.Code = http.StatusOK
		resp.Message = "Updated time: "
		resp.Data = data
		w.WriteHeader(http.StatusOK)

	}
	json.NewEncoder(w).Encode(resp)

}
