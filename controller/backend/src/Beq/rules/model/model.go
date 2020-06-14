package model

import "net/http"

//RulesDataRow is used in DataBase
type RulesDataRow struct {
	FlowID    string
	Protocol  string
	NodeID    string
	DstIP     string
	Interface string
	DstMAC    string
	NodeIP    string
	IsSet     bool
}

//Rule is used for add a rule for Node
type Rule struct {
	RuleID    string `json:"RuleId"`
	Protocol  string `json:"Protocol"`
	FlowID    string `json:"FlowId"`
	DstIP     string `json:"DstIP"`
	Interface string `json:"Interface"`
	DstMAC    string `json:"DstMAC"`
	NodeIP    string `json:"NodeIP"`
}

//Populate populating the data
func (obj *Rule) Populate(RuleID string, data RulesDataRow) {
	obj.RuleID = RuleID
	obj.DstIP = data.DstIP
	obj.DstMAC = data.DstMAC
	obj.FlowID = data.FlowID
	obj.Interface = data.Interface
	obj.NodeIP = data.NodeIP
}

//StateRequest is used for request node state
type StateRequest struct {
	NodeIP string
}

// Response Used for exchange State
type Response struct {
	Program string
	Version string
	Status  string
	Code    int
	Message string
	Data    map[string]interface{}
}

// Default set default value to the response
func (obj *Response) Default() {
	obj.Program = "Beq"
	obj.Version = "0.01"
}

// BadRequest set as Bad Request
func (obj *Response) BadRequest() {
	obj.Code = http.StatusBadRequest
	obj.Status = "Failed"
	obj.Message = "Bad Request"
}

// InternalServerError set Internal server error  Request
func (obj *Response) InternalServerError() {
	obj.Code = http.StatusInternalServerError
	obj.Status = "Failed"
	obj.Message = "Internal Server Error"
}
