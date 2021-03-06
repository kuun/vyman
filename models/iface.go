package models

const IfaceTypeEthernet = "ethernet"

type Iface struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Mac         string `json:"mac"`
	Description string `json:"description"`
}
