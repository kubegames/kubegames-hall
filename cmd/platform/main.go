package main

import (
	"flag"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	service "github.com/kubegames/kubegames-hall/app/service/platform"
	"github.com/kubegames/kubegames-hall/app/service/platform/docs"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/internal/pkg/server"
	"github.com/kubegames/kubegames-hall/internal/pkg/signals"
	"github.com/kubegames/kubegames-hall/pkg/config"
	"github.com/kubegames/kubegames-hall/pkg/platform"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	help bool
	port int64
	cfg  string
)

func init() {
	flag.Int64Var(&port, "p", 8080, "http server port")
	flag.Int64Var(&port, "port", 8080, "http server port")

	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")

	flag.StringVar(&cfg, "c", "./config/config.ini", "config path")
	flag.StringVar(&cfg, "cfg", "./config/config.ini", "config path")

	flag.Parse()
}

func main() {
	//flag
	if help {
		flag.Usage()
		return
	}

	//new config
	config := config.NewConfig(port, cfg)

	//new nats
	nats := nats.NewNats(&config.NatsCfg)
	//nsq clsoe
	defer nats.Engine().Close()
	//drain
	defer nats.Engine().Drain()

	//new mysql
	mysql := xorm.NewMysql(config)

	//new nsq
	nsq := nsq.NewNsq(&config.NsqCfg)

	//new redis
	//redis := redis.NewRedis(&config.RedisCfg)

	//new server
	server := server.NewServer()

	//use pprof
	pprof.Register(server.HttpServer())

	//register swagger docs
	server.HttpServer().GET("/platform/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/swagger.json" {
			c.String(200, docs.GetDocs())
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
			c.URL = "swagger.json"
		})(c)
	})

	platform := platform.NewService(mysql, nats, nsq)
	service.RegisterPlatformServiceHTTPServer(server.HttpServer(), platform)
	service.RegisterPlatformServiceServer(server.GrpcServer(), platform)

	//linsten signal
	signals.RunningSignalHandler(func(c <-chan struct{}) {
		//start
		log.Infof("run kubegames platform server :%d", port)
		if err := server.ListenHttpAndGrpcServe(fmt.Sprintf(":%d", port)); err != nil {
			panic(err)
		}
	})
}
