package model

// PacketDetails is used to store data of incomming packets
type PacketDetails struct {
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
}
