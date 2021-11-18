package models

const IfaceTypeEthernet = "ethernet"

type Iface struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Mac         string `json:"mac"`
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	LinkStatus  string `json:"linkStatus"`
}

type IfaceLinkInfo struct {
	IfIndex   uint16   `json:"ifindex"`
	MTU       uint16   `json:"mtu"`
	IfName    string   `json:"ifname"`
	Flags     []string `json:"flags"`
	Qdisc     string   `json:"qdisc"`
	OperState string   `json:"operstate"`
	LinkMode  string   `json:"linkmode"`
	Group     string   `json:"group"`
	TxQlen    uint16   `json:"txqlen"`
	LinkType  string   `json:"link_type"`
	Address   string   `json:"address"`
	Brodcast  string   `json:"broadcast"`
}

func (info *IfaceLinkInfo) IsLinkUp() bool {
	return info.OperState == "UP"
}
