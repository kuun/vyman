package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
	"github.com/kuun/vyman/models"
	"github.com/kuun/vyman/services"
)

type __logger struct{}

var log = slog.GetLogger(__logger{})

func IsDHCP(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()
	isDHCP, err := service.IsDHCP(ifaceType, name)
	if err != nil {
		goto ERROR
	}
	c.JSON(200, gin.H{"isDHCP": isDHCP})
	return
ERROR:
	log.Errorf("handle request failed, error: %+v", err)
	c.JSON(500, gin.H{"success": false, "error": err.Error()})
}

func EnableDHCP(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()
	err := service.EnableDHCP(ifaceType, name)
	if err != nil {
		log.Errorf("failed to enable dhcp on interface, type: %s, name: %s", ifaceType, name)
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func DisableDHCP(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()
	err := service.DisableDHCP(ifaceType, name)
	if err != nil {
		log.Errorf("failed to disable dhcp on interface, type: %s, name: %s, err: %+v", ifaceType, name, err)
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func GetAddress(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()

	addrs, err := service.GetAddress(ifaceType, name)
	if err != nil {
		log.Errorf("failed to get address, type: %s, name: %s, err: %+v", ifaceType, name, err)
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		c.JSON(200, addrs)
	}
}

func AddAddress(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()

	strAddr := c.Query("addr")

	addr, err := models.ParseAddress(strAddr)
	if err != nil {
		goto ERROR
	}
	if err = service.AddAddress(ifaceType, name, addr); err != nil {
		goto ERROR
	}
	c.JSON(200, gin.H{"success": true})
	return
ERROR:
	log.Errorf("failed to add address, type: %s, name: %s, addr: %s, err: %+v", ifaceType, name, strAddr, err)
	c.JSON(500, gin.H{"success": false, "error": err.Error()})
}

func DeleteAddress(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetAddressService()

	strAddr := c.Query("addr")

	addr, err := models.ParseAddress(strAddr)
	if err != nil {
		goto ERROR
	}
	if err = service.DeleteAddress(ifaceType, name, addr); err != nil {
		goto ERROR
	}
	c.JSON(200, gin.H{"success": true})
	return
ERROR:
	log.Errorf("failed to delete address, type: %s, name: %s, addr: %s, err: %+v", ifaceType, name, strAddr, err)
	c.JSON(500, gin.H{"success": false, "error": err.Error()})
}
