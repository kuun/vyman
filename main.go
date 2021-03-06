package main

import (
	"flag"
	"os"

	"github.com/kuun/vyman/logger"
	"github.com/kuun/vyman/services/vyosclient"
	"github.com/kuun/vyman/session"

	"github.com/gin-gonic/gin"

	"github.com/kuun/slog"
	"github.com/kuun/vyman/routers"
)

type _logger struct{}

var log = slog.GetLogger(&_logger{})

var printHelp = flag.Bool("help", false, "print this usage")
var addr = flag.String("addr", "127.0.0.1:8080", "server listen address")

func main() {
	flag.Parse()

	if *printHelp {
		flag.Usage()
		return
	}

	logger.InitLog()
	engine := gin.Default()

	session.InitSession(engine, "/var/run/vyman", "vyman")
	routers.InitRouters(engine)
	vyosclient.InitVyosClient("127.0.0.1", 8443, os.Getenv("HTTP_API_KEY"))

	log.Warnf("vyman is starting, listen on: %s", *addr)
	if err := engine.Run(*addr); err != nil {
		log.Errorf("can't start vyman, error: %s", err)
	}
}
