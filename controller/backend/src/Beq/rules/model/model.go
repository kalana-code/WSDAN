package model

import "net/http"

//ACTION used for define actions
// type ACTION string

// const (
// 	//Accept used for accept a packet
// 	Accept ACTION = "ACCEPT"
// 	//Drop used for drop a packet
// 	Drop = "DROP"
// )

//RulesDataRow is used in DataBase
type RulesDataRow struct {
	FlowID    string
	Protocol  string
	DstIP     string
	SrcIP     string
	Interface string
	DstMAC    string
	NodeIP    string
	NodeName  string
	IsSet     bool
	Action    string
	IsActive  bool `json:"IsActive"`
}

//ChangeState used for toggle state
func (obj *RulesDataRow) ChangeState() bool {
	obj.IsActive = !obj.IsActive
	return obj.IsActive
}

//Rule is used for add a rule for Node
type Rule struct {
	RuleID    string `json:"RuleId"`
	Protocol  string `json:"Protocol"`
	FlowID    string `json:"FlowId"`
	SrcIP     string `json:"SrcIP"`
	DstIP     string `json:"DstIP"`
	Interface string `json:"Interface"`
	DstMAC    string `json:"DstMAC"`
	NodeIP    string `json:"NodeIP"`
	NodeName  string `json:"Name"`
	Action    string `json:"Action"`
	IsActive  bool   `json:"IsActive"`
}

//Populate populating the data
func (obj *Rule) Populate(RuleID string, data RulesDataRow) {
	obj.RuleID = RuleID
	obj.Protocol = data.Protocol
	obj.DstIP = data.DstIP
	obj.SrcIP = data.SrcIP
	obj.DstMAC = data.DstMAC
	obj.FlowID = data.FlowID
	obj.Interface = data.Interface
	obj.NodeIP = data.NodeIP
	obj.Action = data.Action
	obj.NodeName = data.NodeName
	obj.IsActive = data.IsActive
}

//StateRequest is used for request node state
type StateRequest struct {
	NodeIP string
}

//FlowData used for track flow rule src and destination IP
type FlowData struct {
	DstIP    string
	SrcIP    string
	Protocol string
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
