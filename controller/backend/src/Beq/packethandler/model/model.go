package model

// PacketDetails is used to store data of incomming packets
type PacketDetails struct {
	SrcIP    string `json:"SrcIP"`
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
}
