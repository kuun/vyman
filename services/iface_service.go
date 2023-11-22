package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kuun/vyman/models"
	"github.com/kuun/vyman/services/vyosclient"
	"github.com/kuun/vyman/utils"
	"github.com/pkg/errors"
)

type IfaceService interface {
	GetAllIfaces() ([]*models.Iface, error)
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

func (service *ifaceService) GetAllIfaces() ([]*models.Iface, error) {
	ifaces, err := service.getEthIfaces()
	return ifaces, err
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
	iface.Disabled, err = service.ifaceIsDisabled(ifType, name)
	if err != nil {
		log.Errorf("can't get interface administrator state, interface: %s, error: %+v", err)
	}
	linkInfo, err := service.GetIfaceLinkInfo(iface.Name)
	if err != nil {
		log.Errorf("can't get link info for interface %v, error: %+v", iface.Name, err)
	} else {
		iface.LinkStatus = linkInfo.OperState
	}
	log.Debugf("got interface: %#v", iface)
	return iface, nil
}

func (service *ifaceService) getEthIfaces() ([]*models.Iface, error) {
	prefixes := []string{"lan", "eth", "eno", "ens", "enp", "enx"}
	ifaceNames, err := service.getIfaceNamesWithPrefix(prefixes)
	if err != nil {
		return nil, err
	}
	ifaces := make([]*models.Iface, 0, len(ifaceNames))
	for _, ifaceName := range ifaceNames {
		iface, err := service.GetIface(models.IfaceTypeEthernet, ifaceName)
		if err != nil {
			log.Errorf("can't get interface, type: %v, name: %v", models.IfaceTypeEthernet, ifaceName)
		} else {
			ifaces = append(ifaces, iface)
		}
	}
	return ifaces, nil
}

func (service *ifaceService) getIfaceNamesWithPrefix(prefixes []string) ([]string, error) {
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
	ifaces := make([]string, 0)
	for _, file := range files {
		if filter(file.Name(), prefixes) {
			ifaces = append(ifaces, file.Name())
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
	result, err := service.vyosClient.ReturnValue(fmt.Sprintf("interfaces %s %s description", ifaceType, name))
	if err != nil {
		return "", err
	}
	if !result.Success {
		return "", errors.New(result.Error)
	}
	if result.Data == nil {
		return "", nil
	}
	return result.Data.(string), nil
}

func (service *ifaceService) ifaceIsDisabled(ifaceType, name string) (bool, error) {
	value, err := service.vyosClient.ReturnValue(fmt.Sprintf("interface %s %s disable", ifaceType, name))
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	log.Debugf("interface %s disable state: %v", name, value)
	if value.Data == "true" {
		return true, nil
	}
	return false, nil
}

func (service *ifaceService) GetIfaceLinkInfo(ifaceName string) (*models.IfaceLinkInfo, error) {
	output, err := utils.CmdRunOutput("ip", "-j", "link", "show", ifaceName)
	if err != nil {
		return nil, err
	}
	infos := make([]models.IfaceLinkInfo, 0, 1)
	if err = json.Unmarshal([]byte(output), &infos); err != nil {
		return nil, errors.Wrap(err, "")
	}
	if len(infos) != 1 {
		log.Errorf("invalid ip link output: %v", output)
		return nil, errors.New("invalid ip link output")
	}
	return &infos[0], nil
}
