package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/vyman/services/vyosclient"
	"github.com/pkg/errors"
	"strings"
)

func ShowConfig(c *gin.Context) {
	client := vyosclient.GetClient()

	reqData := make(map[string]string)
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Errorf("failed to parse client request, error: %+v", err)
		c.Status(400)
		return
	}
	path, ok := reqData["path"]
	if !ok || path == "" {
		log.Errorf("no path field in request", err)
		c.Status(400)
		return
	}
	result, err := client.ShowConfig(path)
	if err != nil {
		log.Errorf("failed to show config, error: %+v", err)
		c.JSON(200, gin.H{"success": false, "error": errors.Cause(err).Error()})
		return
	}
	if !result.Success && strings.Contains(result.Error, "Configuration under specified path is empty") {
		result.Success = true
	}
	c.JSON(200, result)
}

func ReturnValue(c *gin.Context) {
	client := vyosclient.GetClient()

	reqData := make(map[string]string)
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Errorf("failed to parse client request, error: %+v", err)
		c.Status(400)
		return
	}
	path, ok := reqData["path"]
	if !ok || path == "" {
		log.Errorf("no path field in request", err)
		c.Status(400)
		return
	}
	result, err := client.ReturnValue(path)
	if err != nil {
		log.Errorf("failed to return result, error: %+v", err)
		c.JSON(200, gin.H{"success": false, "error": errors.Cause(err)})
		return
	}
	c.JSON(200, result)
}

func ReturnValues(c *gin.Context) {
	client := vyosclient.GetClient()

	reqData := make(map[string]string)
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Errorf("failed to parse client request, error: %+v", err)
		c.Status(400)
		return
	}
	path, ok := reqData["path"]
	if !ok || path == "" {
		log.Errorf("no path field in request", err)
		c.Status(400)
		return
	}
	result, err := client.ReturnValues(path)
	if err != nil {
		log.Errorf("failed to return result, error: %+v", err)
		c.JSON(200, gin.H{"success": false, "error": errors.Cause(err)})
		return
	}
	c.JSON(200, result)
}

func Configure(c *gin.Context) {
	client := vyosclient.GetClient()

	cmds := make([]*vyosclient.Command, 0)
	err := c.BindJSON(&cmds)
	if err != nil {
		log.Errorf("failed to parse client request, error: %+v", err)
		c.Status(400)
		return
	}

	err = client.Configure(cmds)
	if err != nil {
		log.Errorf("failed to execute commands, cmds: %#v, error: %+v", cmds, err)
		c.JSON(200, gin.H{"success": false, "error": errors.Cause(err)})
		return
	}
	c.JSON(200, gin.H{"success": true})
}
