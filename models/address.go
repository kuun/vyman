package models

import "fmt"

type Address struct {
	Local             string `json:"local"`
	PrefixLen         uint16 `json:"prefixlen"`
	Dynamic           bool   `json:"dynamic"`
	Family            string `json:"family"`
	Broadcast         string `json:"broadcast"`
	Scope             string `json:"scope"`
	ValidLifeTime     uint64 `json:"valid_life_time"`
	PreferredLifeTime uint64 `json:"preferred_life_time"`
}

type AddrInfoWrapper struct {
	AddrInfos []Address `json:"addr_info"`
}

func (addr *Address) CIDR() string {
	return fmt.Sprintf("%s/%d", addr.Local, addr.PrefixLen)
}
