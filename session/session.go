package session

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/kuun/slog"
)

var store *sessions.FilesystemStore

const sessionName = "vyman"

type _logger struct{}

var log = slog.GetLogger(&_logger{})

func InitSession(engine *gin.Engine, sessionPath, sessionKey string) {
	store = sessions.NewFilesystemStore(sessionPath, []byte(sessionKey))
	store.MaxAge(3600)

	engine.Use(authFilter())
}

func authFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if strings.HasSuffix(uri, ".js") {
			return
		}
		if strings.HasSuffix(uri, ".css") {
			return
		}
		if strings.HasSuffix(uri, ".png") {
			return
		}
		if uri == "/login.html" {
			return
		}
		if uri == "/api/login" {
			return
		}
		session, err := store.New(c.Request, sessionName)
		if err != nil {
			log.Errorf("can't create session, error: %v", err)
		}
		if session.IsNew {
			c.Redirect(302, "/login.html")
			c.Abort()
		}
	}

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	session, _ := store.Get(c.Request, sessionName)
	if username == os.Getenv("VYMAN_USER") && password == os.Getenv("VYMAN_PASSWORD") {
		session.AddFlash("username", username)
		err := session.Save(c.Request, c.Writer)
		if err != nil {
			log.Errorf("can't save session, error: %s", err)
		}
		c.JSON(200, gin.H{"success": true})
	} else {
		log.Warnf("user %s login failed", username)
		session.Options.MaxAge = 0
		session.Save(c.Request, c.Writer)
		c.JSON(403, gin.H{"success": false})
	}
}

func Logout(c *gin.Context) {
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		log.Errorf("can't create session, error: %v", err)
		c.JSON(500, gin.H{"success": false})
		return
	}
	session.Options.MaxAge = 0
	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{"success": true})
}
