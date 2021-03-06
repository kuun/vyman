package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/vyman/services"
)

func GetIfaces(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	service := services.GetIfaceService()

	ifaces, err := service.GetIfaces(ifaceType)
	if err != nil {
		log.Errorf("failed to get all ifaces, type: %s, err: %+v", ifaceType, err)
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		c.JSON(200, ifaces)
	}
}

func GetIface(c *gin.Context) {
	ifaceType := c.Param("ifaceType")
	name := c.Param("name")
	service := services.GetIfaceService()

	ifaces, err := service.GetIface(ifaceType, name)
	if err != nil {
		log.Errorf("failed to get all ifaces, type: %s, err: %+v", ifaceType, err)
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
	} else {
		c.JSON(200, ifaces)
	}
}
