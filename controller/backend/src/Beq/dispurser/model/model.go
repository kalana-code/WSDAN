package model

//JobType used for define job type
type JobType int

const (
	//TypeAddRule used for identify rule set task
	TypeAddRule JobType = 1

	//TypeRemoveRule job type used for create job for remove a rule in node
	TypeRemoveRule = 2

	//TypeNodeState used for request  node state task
	TypeNodeState = 3
)

const (
	//Port used to define Port of rule endpoint
	Port = "8082"
	//RemoveRuleEndPoint used to define remove rule endpoint
	RemoveRuleEndPoint = "endpoint"
	//AddRuleEndPoint used to define add rule endpoint
	AddRuleEndPoint = "AddFlowRule"
)

//Job is used for add Tesk for task queue
type Job struct {
	Type        JobType
	NodeIP      string
	TaskDetails interface{}
}

//RemoveRuleJob used for remove rule from node
type RemoveRuleJob struct {
	RuleID string
}

//AddRuleJob used for add rule to node
type AddRuleJob struct {
	RuleID    string `json:"RuleId"`
	Protocol  string `json:"Protocol"`
	FlowID    string `json:"FlowId"`
	DstIP     string `json:"DstIP"`
	Interface string `json:"Interface"`
	DstMAC    string `json:"DstMAC"`
	Action    string `json:"DstMAC"`
}
