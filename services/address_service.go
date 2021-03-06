package services

import (
	"errors"
	"fmt"

	"github.com/kuun/slog"
	"github.com/kuun/vyman/models"
	"github.com/kuun/vyman/services/vyosclient"
)

var log = slog.GetLogger(addressService{})

type AddressService interface {
	IsDHCP(ifaceType string, name string) (bool, error)
	EnableDHCP(ifaceType string, name string) error
	DisableDHCP(ifaceType string, name string) error
	GetAddress(ifaceType string, name string) ([]*models.Address, error)
	AddAddress(ifaceType string, name string, addr *models.Address) error
	DeleteAddress(ifaceType string, name string, addr *models.Address) error
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

func (service *addressService) GetAddress(ifaceType string, name string) ([]*models.Address, error) {
	addrs := make([]*models.Address, 0)
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	values, err := service.vyosClient.ReturnValues(path)
	if err != nil {
		return addrs, err
	}
	for _, value := range values {
		if value == "dhcp" {
			return addrs[:0], nil
		}
		addr, err := models.ParseAddress(value)
		if err != nil {
			log.Errorf("failed to parse ip address, value: %s, error: %+v", value, err)
		} else {
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}

func (service *addressService) AddAddress(ifaceType string, name string, addr *models.Address) error {
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := []*vyosclient.Command{
		vyosclient.NewCommand("set", path, addr.String()),
	}
	return service.vyosClient.Configure(cmds)
}

func (service *addressService) DeleteAddress(ifaceType string, name string, addr *models.Address) error {
	path := fmt.Sprintf("interfaces %s %s address", ifaceType, name)
	cmds := []*vyosclient.Command{
		vyosclient.NewCommand("delete", path, addr.String()),
	}
	return service.vyosClient.Configure(cmds)
}
