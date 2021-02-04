package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
	"github.com/kuun/vyman/session"
)

type _logger struct{}

var log = slog.GetLogger(_logger{})

// InitRouters setups routers
func InitRouters(r *gin.Engine) {
	r.Static("/thirdpart", "webapp/thirdpart")
	r.Static("/css", "webapp/css")
	r.Static("/dist", "webapp/dist")
	r.StaticFile("/", "webapp/index.html")
	r.StaticFile("/index.html", "webapp/index.html")
	r.StaticFile("/login.html", "webapp/login.html")

	r.POST("/api/login", session.Login)
	r.POST("/api/logout", session.Logout)
}
