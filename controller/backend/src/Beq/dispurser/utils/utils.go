package utils

// called endpoint service in node base on type of job as following

import (
	"Beq/dispurser/model"
	"bytes"
	"log"
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
	// host := net.ParseIP(job.NodeIP)
	// if host == nil {
	// 	log.Println("ERROR: [RD]: [Remove Rule] IP convertion process was failed")
	// 	return
	// }

	//figureout remove rule model
	data, ok := job.TaskDetails.(model.RemoveRuleJob)
	if !ok {
		log.Println("ERROR: [RD]: [Remove Rule] Task details not match to RemoveRuleJob model")
		return
	}
	httpClient(job.NodeIP, model.RemoveRuleEndPoint, data)

}

func addRule(job *model.Job) {
	// figureout host IP
	// host := net.ParseIP(job.NodeIP)
	// if host == nil {
	// 	log.Println("ERROR: [RD]: [Remove Rule] IP convertion process was failed")
	// 	return
	// }

	//figureout remove rule model
	// data, ok := job.TaskDetails.(model.AddRuleJob)
	// if !ok {
	// 	log.Println("ERROR: [RD]: [Add Rule] Task details not match to AddRuleJob model")
	// 	return
	// }
	httpClient(job.NodeIP, model.AddRuleEndPoint, job.TaskDetails)
}

func httpClient(host string, endPoint string, data interface{}) {
	url := host + ":" + model.Port + "/" + endPoint
	jsonData, ok := data.([]byte)
	if !ok {
		log.Println("ERROR: [HC]: [HTTP Client] Request data convertion [to []byte ] process was failed")
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
}
