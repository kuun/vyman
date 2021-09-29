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
	r.Static("/thirdpart", "webapp/thirdpart")
	r.Static("/css", "webapp/css")
	r.Static("/dist", "webapp/dist")
	r.Static("/views", "webapp/src/views")
	r.StaticFile("/", "webapp/index.html")
	r.StaticFile("/index.html", "webapp/index.html")
	r.StaticFile("/login.html", "webapp/login.html")

	r.POST("/api/login", session.Login)
	r.POST("/api/logout", session.Logout)

	r.GET("/api/interfaces/:ifaceType", handlers.GetIfaces)
	r.GET("/api/interfaces/:ifaceType/:name", handlers.GetIface)

	r.GET("/api/interfaces/:ifaceType/:name/dhcp/enabled", handlers.IsDHCP)
	r.POST("/api/interfaces/:ifaceType/:name/dhcp/enable", handlers.EnableDHCP)
	r.POST("/api/interfaces/:ifaceType/:name/dhcp/disable", handlers.DisableDHCP)
	r.GET("/api/interfaces/:ifaceType/:name/address", handlers.GetAddress)
	r.POST("/api/interfaces/:ifaceType/:name/address", handlers.AddAddress)
	r.DELETE("/api/interfaces/:ifaceType/:name/address", handlers.DeleteAddress)

}
