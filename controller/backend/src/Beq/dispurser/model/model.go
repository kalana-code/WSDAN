package model

//JobType used for define job type
type JobType int

const (
	//RuleDispurse used for identify rule set task
	RuleDispurse JobType = 1
	//NodeState used for request  node state task
	NodeState = 2
)

//Job is used for add Tesk for task queue
type Job struct {
	Type        JobType
	IP          string
	TaskDetails interface{}
}


