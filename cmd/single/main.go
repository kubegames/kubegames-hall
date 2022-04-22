package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/kubegames/kubegames-hall/app/service/docs"
	gameService "github.com/kubegames/kubegames-hall/app/service/game"
	platformService "github.com/kubegames/kubegames-hall/app/service/platform"
	playerService "github.com/kubegames/kubegames-hall/app/service/player"
	roomService "github.com/kubegames/kubegames-hall/app/service/room"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/internal/pkg/server"
	"github.com/kubegames/kubegames-hall/internal/pkg/signals"
	"github.com/kubegames/kubegames-hall/pkg/config"
	"github.com/kubegames/kubegames-hall/pkg/game"
	"github.com/kubegames/kubegames-hall/pkg/platform"
	"github.com/kubegames/kubegames-hall/pkg/player"
	"github.com/kubegames/kubegames-hall/pkg/room"
	gamesclientset "github.com/kubegames/kubegames-operator/pkg/client/game/clientset/versioned"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	help       bool
	port       int64
	cfg        string
	kubeconfig string
)

func init() {
	flag.Int64Var(&port, "p", 8080, "http server port")
	flag.Int64Var(&port, "port", 8080, "http server port")

	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")

	flag.StringVar(&cfg, "c", "./config/config.ini", "config path")
	flag.StringVar(&cfg, "cfg", "./config/config.ini", "config path")

	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) kubeconfig absolute path to the file")
		flag.StringVar(&kubeconfig, "k", filepath.Join(home, ".kube", "config"), "(optional) kubeconfig absolute path to the file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "(optional) kubeconfig absolute path to the file")
		flag.StringVar(&kubeconfig, "k", "", "(optional) kubeconfig absolute path to the file")
	}

	flag.Parse()
}

func main() {
	//flag
	if help {
		flag.Usage()
		return
	}

	//new k8s client
	var k8sconfig *rest.Config
	var err error

	if len(kubeconfig) > 0 {
		if k8sconfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig); err != nil {
			panic(err.Error())
		}
	} else {
		if k8sconfig, err = rest.InClusterConfig(); err != nil {
			panic(err.Error())
		}
	}

	k8sclient, err := kubernetes.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}

	gamesclientset, err := gamesclientset.NewForConfig(k8sconfig)
	if err != nil {
		panic(err)
	}

	//new config
	config := config.NewConfig(port, cfg)

	//new redis
	//redis := redis.NewRedis(&config.RedisCfg)

	//new nats
	nats := nats.NewNats(&config.NatsCfg)
	//nsq clsoe
	defer nats.Engine().Close()
	//drain
	defer nats.Engine().Drain()

	//new nsq
	nsq := nsq.NewNsq(&config.NsqCfg)

	//new mysql
	mysql := xorm.NewMysql(config)

	//new server
	server := server.NewServer()

	//use pprof
	pprof.Register(server.HttpServer())

	//register swagger docs
	server.HttpServer().GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/swagger.json" {
			c.String(200, docs.GetDocs())
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
			c.URL = "swagger.json"
		})(c)
	})

	//register player
	player := player.NewService(mysql, nats)
	server.HttpServer().GET("/player/websocket", player.Websocket)
	playerService.RegisterPlayerServiceHTTPServer(server.HttpServer(), player)
	playerService.RegisterPlayerServiceServer(server.GrpcServer(), player)

	//register platform
	platform := platform.NewService(mysql, nats, nsq)
	platformService.RegisterPlatformServiceHTTPServer(server.HttpServer(), platform)
	platformService.RegisterPlatformServiceServer(server.GrpcServer(), platform)

	//register room
	room := room.NewService(mysql, nsq)
	roomService.RegisterRoomServiceHTTPServer(server.HttpServer(), room)
	roomService.RegisterRoomServiceServer(server.GrpcServer(), room)

	//register game
	game := game.NewService(mysql, k8sclient, gamesclientset)
	gameService.RegisterGameServiceHTTPServer(server.HttpServer(), game)
	gameService.RegisterGameServiceServer(server.GrpcServer(), game)

	//linsten signal
	signals.RunningSignalHandler(func(c <-chan struct{}) {
		//start
		log.Infof("run kubegames player server :%d", port)
		if err := server.ListenHttpAndGrpcServe(fmt.Sprintf(":%d", port)); err != nil {
			panic(err)
		}
	})

}
