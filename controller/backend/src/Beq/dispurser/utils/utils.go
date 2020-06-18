package utils

// called endpoint service in node base on type of job as following

import (
	"Beq/dispurser/model"
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
)

//Dispurse used for dispurse jobs to Node
func Dispurse(job *model.Job) {
	log.Println("INFO: [RD]: Dispursing .... ")
	switch job.Type {
	case model.TypeAddRule:
		// add rule
		addRule(job)
		break
	case model.TypeRemoveRule:
		// remove rule
		removeRule(job)
		break
	case 3:
		// add flow
		break
	case 4:
		// remove flow
		break
	}
}

func removeRule(job *model.Job) {
	// figureout host IP
	host := net.ParseIP(job.NodeIP)
	if host == nil {
		log.Println("ERROR: [RD]: [Remove Rule] IP convertion process was failed")
		return
	}

	//figureout remove rule model
	data, ok := job.TaskDetails.(model.RemoveRuleJob)
	if !ok {
		log.Println("ERROR: [RD]: [Remove Rule] Task details not match to RemoveRuleJob model")
		return
	}
	httpClient(host, model.RemoveRuleEndPoint, data)

}

func addRule(job *model.Job) {
	// figureout host IP
	host := net.ParseIP(job.NodeIP)
	if host == nil {
		log.Println("ERROR: [RD]: [Remove Rule] IP convertion process was failed")
		return
	}

	// figureout remove rule model
	data, ok := job.TaskDetails.(model.AddRuleJob)
	if !ok {
		log.Println("ERROR: [RD]: [Add Rule] Task details not match to AddRuleJob model")
		return
	}
	httpClient(host, model.AddRuleEndPoint, data)
}

func httpClient(host net.IP, endPoint string, data interface{}) {

	url := "http://" + host.String() + ":" + model.Port + "/" + endPoint
	AddRule, isAddRule := data.(model.AddRuleJob)
	RemoveJob, isRemoveJob := data.(model.RemoveRuleJob)
	if isAddRule {
		jsonData := map[string]interface{}{
			"RuleID":    AddRule.RuleID,
			"Protocol":  AddRule.Protocol,
			"FlowID":    AddRule.FlowID,
			"DstIP":     AddRule.DstIP,
			"Interface": AddRule.Interface,
			"DstMAC":    AddRule.DstMAC,
			"Action":    AddRule.Action,
		}
		jsonValue, _ := json.Marshal(jsonData)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("ERROR: [HC]: [HTTP Client] Request failed.")
			return
		}
		defer resp.Body.Close()
		log.Println("INFO: [HC]: [HTTP Client] Status Request.", resp.Status)
	} else if isRemoveJob {
		jsonData := map[string]interface{}{
			"RuleID": RemoveJob.RuleID,
		}
		jsonValue, _ := json.Marshal(jsonData)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("ERROR: [HC]: [HTTP Client] Request failed.")
			return
		}
		defer resp.Body.Close()
		log.Println("INFO: [HC]: [HTTP Client] Status Request.", resp.Status)
	} else {
		log.Println("INFO: [HC]: Invalid JOB data")
	}

}
