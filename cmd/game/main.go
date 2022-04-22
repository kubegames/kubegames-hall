package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	service "github.com/kubegames/kubegames-hall/app/service/game"
	"github.com/kubegames/kubegames-hall/app/service/game/docs"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/internal/pkg/server"
	"github.com/kubegames/kubegames-hall/internal/pkg/signals"
	"github.com/kubegames/kubegames-hall/pkg/config"
	"github.com/kubegames/kubegames-hall/pkg/game"
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

	//new mysql
	mysql := xorm.NewMysql(config)

	//new server
	server := server.NewServer()

	//use pprof
	pprof.Register(server.HttpServer())

	//register swagger docs
	server.HttpServer().GET("/game/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/swagger.json" {
			c.String(200, docs.GetDocs())
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
			c.URL = "swagger.json"
		})(c)
	})

	game := game.NewService(mysql, k8sclient, gamesclientset)
	service.RegisterGameServiceHTTPServer(server.HttpServer(), game)
	service.RegisterGameServiceServer(server.GrpcServer(), game)

	//linsten signal
	signals.RunningSignalHandler(func(c <-chan struct{}) {
		//start
		log.Infof("run kubegames game server :%d", port)
		if err := server.ListenHttpAndGrpcServe(fmt.Sprintf(":%d", port)); err != nil {
			panic(err)
		}
	})
}
