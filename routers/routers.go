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
	r.Static("/thirdpart", "static/thirdpart")
	r.Static("/css", "static/css")
	r.StaticFile("/index.html", "static/index.html")
	r.StaticFile("/login.html", "static/login.html")

	r.POST("/api/login", session.Login)
	r.POST("/api/logout", session.Logout)
}
