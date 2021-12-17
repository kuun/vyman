package services

import (
	"encoding/json"
	"fmt"

	"github.com/kuun/slog"
	"github.com/kuun/vyman/models"
	"github.com/kuun/vyman/services/vyosclient"
	"github.com/kuun/vyman/utils"
	"github.com/pkg/errors"
)

var log = slog.GetLogger(addressService{})

type AddressService interface {
	IsDHCP(ifaceType string, name string) (bool, error)
	EnableDHCP(ifaceType string, name string) error
	DisableDHCP(ifaceType string, name string) error
	GetAddress(ifaceType string, name string) ([]models.Address, error)
	AddAddress(ifaceType string, name string, addr *models.Address) error
	DeleteAddress(ifaceType string, name string, addrs []*models.Address) error
}

type addressService struct {
	vyosClient vyosclient.VyosClient
}

func GetAddressService() AddressService {
	return &addressService{
		vyosClient: vyosclient.GetClient(),
	}
}

func (service *addressService) IsDHCP(ifaceType string, name string) (bool, error) {
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	value, err := service.vyosClient.ReturnValue(path)
	if err != nil {
		return false, err
	}
	if value == "dhcp" {
		return true, nil
	} else {
		return false, nil
	}
}

func (service *addressService) EnableDHCP(ifaceType string, name string) error {
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := []*vyosclient.Command{
		vyosclient.NewCommand("delete", path, ""),
		vyosclient.NewCommand("set", path, "dhcp"),
	}
	return service.vyosClient.Configure(cmds)
}

func (service *addressService) DisableDHCP(ifaceType string, name string) error {
	if isDHCP, err := service.IsDHCP(ifaceType, name); err != nil {
		return err
	} else if !isDHCP {
		log.Warnf("current interface is not dhcp enabled, interface type: %s, name: %s", ifaceType, name)
		return errors.New("该网卡未启用DHCP")
	}
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := []*vyosclient.Command{
		vyosclient.NewCommand("delete", path, "dhcp"),
	}
	return service.vyosClient.Configure(cmds)
}

func (service *addressService) GetAddress(ifaceType string, name string) ([]models.Address, error) {
	output, err := utils.CmdRunOutput("ip", "-j", "address", "show", "dev", name)
	if err != nil {
		return nil, err
	}
	addrInfoWrapper := make([]models.AddrInfoWrapper, 0, 1)
	if err = json.Unmarshal([]byte(output), &addrInfoWrapper); err != nil {
		return nil, errors.Wrap(err, "")
	}
	if len(addrInfoWrapper) != 1 {
		log.Errorf("invalid ip -j address show dev %s output, output: %s", name, output)
		return nil, errors.New("invalid ip -j address show output")
	}
	return addrInfoWrapper[0].AddrInfos, nil
}

func (service *addressService) AddAddress(ifaceType string, name string, addr *models.Address) error {
	log.Infof("try to add address: %#v, interface type: %s, interface: %s", addr, ifaceType, name)
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := []*vyosclient.Command{
		vyosclient.NewCommand("set", path, addr.CIDR()),
	}
	return service.vyosClient.Configure(cmds)
}

func (service *addressService) DeleteAddress(ifaceType string, name string, addrs []*models.Address) error {
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := make([]*vyosclient.Command, 0, len(addrs))
	for _, addr := range addrs {
		cmds = append(cmds, vyosclient.NewCommand("delete", path, addr.CIDR()))
	}
	return service.vyosClient.Configure(cmds)
}
