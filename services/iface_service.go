package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kuun/vyman/models"
	"github.com/kuun/vyman/services/vyosclient"
	"github.com/pkg/errors"
)

type IfaceService interface {
	GetIfaces(ifType string) ([]*models.Iface, error)
	GetIface(ifType string, name string) (*models.Iface, error)
}

type ifaceService struct {
	vyosClient vyosclient.VyosClient
}

func GetIfaceService() IfaceService {
	return &ifaceService{
		vyosClient: vyosclient.GetClient(),
	}
}

func (service *ifaceService) GetIfaces(ifType string) ([]*models.Iface, error) {
	switch ifType {
	case "ethernet":
		return service.getEthIfaces()
	default:
		return make([]*models.Iface, 0), nil
	}
}

func (service *ifaceService) GetIface(ifType string, name string) (*models.Iface, error) {
	_, err := os.Stat(fmt.Sprintf("/sys/class/net/%s", name))
	if err != nil {
		log.Errorf("interface %s may not be exist, error: %s", name, err)
		return nil, errors.New("网卡不存在")
	}
	iface := &models.Iface{Name: name, Type: ifType}
	iface.Mac, err = service.getMac(iface.Name)
	if err != nil {
		log.Errorf("can't get mac of interface %s, error: %+v", iface.Name, err)
	}
	iface.Description, err = service.getIfaceDescription(iface.Type, iface.Name)
	if err != nil {
		log.Errorf("can't get description of interface %s, error: %+v", iface.Name, err)
	}
	return iface, nil
}

func (service *ifaceService) getEthIfaces() ([]*models.Iface, error) {
	prefixes := []string{"lan", "eth", "eno", "ens", "enp", "enx"}
	ifaces, err := service.getIfacesWithPrefix(prefixes)
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		iface.Type = models.IfaceTypeEthernet
		iface.Mac, err = service.getMac(iface.Name)
		if err != nil {
			log.Errorf("can't get mac of interface %s, error: %+v", iface.Name, err)
		}
		iface.Description, err = service.getIfaceDescription(iface.Type, iface.Name)
		if err != nil {
			log.Errorf("can't get description of interface %s, error: %+v", iface.Name, err)
		}
	}
	return ifaces, nil
}

func (service *ifaceService) getIfacesWithPrefix(prefixes []string) ([]*models.Iface, error) {
	files, err := ioutil.ReadDir("/sys/class/net/")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	filter := func(name string, prefixes []string) bool {
		for _, prefix := range prefixes {
			if strings.HasPrefix(name, prefix) {
				return true
			}
		}
		return false
	}
	ifaces := make([]*models.Iface, 0)
	for _, file := range files {
		if filter(file.Name(), prefixes) {
			ifaces = append(ifaces, &models.Iface{Name: file.Name()})
		}
	}
	return ifaces, nil
}

func (service *ifaceService) getMac(name string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/sys/class/net/%s/address", name))
	if err != nil {
		return "", errors.WithStack(err)
	}
	length := len(content)
	if length <= 0 {
		return "", errors.New("invalid mac")
	}
	if content[length-1] == '\n' {
		content = content[:length-1]
	}
	return string(content), nil
}

func (service *ifaceService) getIfaceDescription(ifaceType string, name string) (string, error) {
	return service.vyosClient.ReturnValue(fmt.Sprintf("interfaces %s %s", ifaceType, name))
}
