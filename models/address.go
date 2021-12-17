package models

import "fmt"

type Address struct {
	Local             string `json:"local" form:"local"`
	PrefixLen         uint16 `json:"prefixlen" form:"prefixlen"`
	Dynamic           bool   `json:"dynamic" form:"dynamic"`
	Family            string `json:"family" form:"family"`
	Broadcast         string `json:"broadcast" form:"broadcast"`
	Scope             string `json:"scope" form:"scope"`
	ValidLifeTime     uint64 `json:"valid_life_time" form:"valid_life_time"`
	PreferredLifeTime uint64 `json:"preferred_life_time" from:"preferred_life_time"`
}

type AddrInfoWrapper struct {
	AddrInfos []Address `json:"addr_info"`
}

func (addr *Address) CIDR() string {
	return fmt.Sprintf("%s/%d", addr.Local, addr.PrefixLen)
}
