package model

import "net/http"

// GrpNode used for geanarate Node information for network graph
type GrpNode struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Group string `json:"group"`
}

// GrpNodeLink  used for keep ling information between two Nodes
type GrpNodeLink struct {
	From               int    `json:"from"`
	To                 int    `json:"to"`
	Label              string `json:"label"`
	arrowStrikethrough bool
	length             int
	dashes             bool
}

func (obj *GrpNodeLink) SetLink() {
	obj.arrowStrikethrough = true
	obj.length = 150
	obj.dashes = true
}

// GrpData keeps network map
type GrpData struct {
	Nodes []GrpNode     `json:"nodes"`
	Edges []GrpNodeLink `json:"edges"`
}

//NodeData used for store data into db
type NodeData struct {
	Node       Node
	Neighbours []Neighbour
}

//Node Inforamation track
type Node struct {
	Name  string
	Group string
	IP    string
	MAC   string
}

// Neighbour struct
type Neighbour struct {
	MAC       string
	Bandwidth int
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
