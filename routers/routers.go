package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
	"github.com/kuun/vyman/handlers"
	"github.com/kuun/vyman/session"
)

type _logger struct{}

var log = slog.GetLogger(_logger{})

// InitRouters setups routers
func InitRouters(r *gin.Engine) {
	r.Use(disableCache())

	r.Static("/assets", "webapp/assets")
	r.StaticFile("/", "webapp/index.html")
	r.StaticFile("/index.html", "webapp/index.html")
	r.GET("/ui/*filepath", func(c *gin.Context) {
		c.File("webapp/index.html")
	})

	r.POST("/api/login", session.Login)
	r.POST("/api/logout", session.Logout)

	r.POST("/api/showConfig", handlers.ShowConfig)
	r.POST("/api/configure", handlers.Configure)
	r.POST("/api/returnValue", handlers.ReturnValue)
	r.POST("/api/returnValues", handlers.ReturnValues)

	r.GET("/api/interfaces", handlers.GetAllIfaces)
	r.GET("/api/interfaces/:ifaceType", handlers.GetIfaces)
	r.GET("/api/interfaces/:ifaceType/:name", handlers.GetIface)

	r.GET("/api/interfaces/:ifaceType/:name/dhcp/enabled", handlers.IsDHCP)
	r.POST("/api/interfaces/:ifaceType/:name/dhcp/enable", handlers.EnableDHCP)
	r.POST("/api/interfaces/:ifaceType/:name/dhcp/disable", handlers.DisableDHCP)
	r.GET("/api/interfaces/:ifaceType/:name/address", handlers.GetAddress)
	r.POST("/api/interfaces/:ifaceType/:name/address", handlers.AddAddress)
	r.DELETE("/api/interfaces/:ifaceType/:name/address", handlers.DeleteAddress)

}

func disableCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")
		c.Next()
	}
}
