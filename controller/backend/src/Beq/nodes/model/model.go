package model

import "net/http"

// GrpNode used for geanarate Node information for network graph
type GrpNode struct {
	ID       int      `json:"id"`
	Label    string   `json:"label"`
	Group    string   `json:"group"`
	NodeData NodeData `json:"NodeData"`
}

// GrpNodeLink  used for keep ling information between two Nodes
type GrpNodeLink struct {
	From               int    `json:"from"`
	To                 int    `json:"to"`
	Label              string `json:"label"`
	ArrowStrikethrough bool   `json:"arrowStrikethrough"`
	Length             int    `json:"length"`
	Dashes             bool   `json:"dashes"`
	Arrows             Arrow  `json:"arrows"`
}

//SetLink used for set default value in edge
func (obj *GrpNodeLink) SetLink(From int, To int, Label string) {
	obj.From = From
	obj.To = To
	obj.Label = Label
	obj.ArrowStrikethrough = false
	obj.Length = 200
	obj.Dashes = true
	// set Arrow head
	obj.Arrows = Arrow{}
	obj.Arrows.To = ArrowStyle{Enabled: false}
	obj.Arrows.From = ArrowStyle{Enabled: false}
	obj.Arrows.Middle = ArrowStyle{Enabled: false}

}

// Arrow keep arrow header style
type Arrow struct {
	To     ArrowStyle `json:"to"`
	Middle ArrowStyle `json:"middle"`
	From   ArrowStyle `json:"from"`
}

// ArrowStyle keep Arrow style
type ArrowStyle struct {
	Enabled bool `json:"enabled"`
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
	Bandwidth string
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
