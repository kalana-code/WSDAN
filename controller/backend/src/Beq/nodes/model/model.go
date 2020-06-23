package model

import "net/http"

//FlowLink used for build Flow Link
type FlowLink struct {
	SrcNode int
	DstNode int
}

// GrpNode used for geanarate Node information for network graph
type GrpNode struct {
	ID       int      `json:"id"`
	Label    string   `json:"label"`
	Group    string   `json:"group"`
	NodeData NodeData `json:"NodeData"`
}

// GrpNodeLink  used for keep ling information between two Nodes
type GrpNodeLink struct {
	From               int         `json:"from"`
	To                 int         `json:"to"`
	Label              string      `json:"label"`
	ArrowStrikethrough bool        `json:"arrowStrikethrough"`
	Color              ColourStyle `json:"color"`
	Dashes             bool        `json:"dashes"`
	Arrows             Arrow       `json:"arrows"`
	Length             int         `json:"length"`
	Physics            bool        `json:"physics"`
}

//SetLink used for set default value in edge
func (obj *GrpNodeLink) SetLink(From int, To int, Label string, controllerLink bool) {
	obj.From = From
	obj.To = To
	obj.Length = 200
	if !controllerLink {
		obj.Label = Label
		// obj.Length = 400
	}
	obj.ArrowStrikethrough = false
	obj.Physics = true
	obj.Dashes = true
	// set Arrow head
	obj.Arrows = Arrow{}
	obj.Color = ColourStyle{
		Color: "gray",
	}
	obj.Arrows.To = ArrowStyle{Enabled: false}
	obj.Arrows.From = ArrowStyle{Enabled: false}
	obj.Arrows.Middle = ArrowStyle{Enabled: false}

}

//SetLinkWithColor used for set color to edge
func (obj *GrpNodeLink) SetLinkWithColor(From int, To int, Label string, controllerLink bool, color string, isDashed bool) {
	obj.From = From
	obj.To = To
	obj.Length = 200
	if !controllerLink {
		obj.Label = Label
		// obj.Length = 400
	}
	obj.ArrowStrikethrough = false
	obj.Physics = true
	obj.Dashes = isDashed
	// set Arrow head
	obj.Arrows = Arrow{}
	obj.Color = ColourStyle{
		Color: color,
	}
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

// ColourStyle keep color style
type ColourStyle struct {
	Color string `json:"color"`
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
