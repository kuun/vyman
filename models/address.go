package models

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

type Address struct {
	Addr string `json:"addr"`
	Mask int    `json:"mask"`
}

func (addr *Address) String() string {
	return fmt.Sprintf("%s/%d", addr.Addr, addr.Mask)
}

func ParseAddress(str string) (*Address, error) {
	ip, ipnet, err := net.ParseCIDR(str)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	addr := &Address{}
	addr.Addr = ip.String()
	ones, _ := ipnet.Mask.Size()
	addr.Mask = ones
	return addr, nil
}
