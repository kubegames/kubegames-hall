package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kubegames/kubegames-hall/internal/pkg/config/ini"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/internal/pkg/redis"
)

//cofnig object
type Config struct {
	Port       int64
	ReadMysql  xorm.MysqlCfg  `ini:"ReadMysql"`
	WriteMysql xorm.MysqlCfg  `ini:"WriteMysql"`
	NsqCfg     nsq.NsqCfg     `ini:"NsqCfg"`
	NatsCfg    nats.NatsCfg   `ini:"NatsCfg"`
	RedisCfg   redis.RedisCfg `ini:"RedisCfg"`
}

//get read mysql
func (c *Config) GetReadMysql() *xorm.MysqlCfg {
	return &c.ReadMysql
}

//get write mysql
func (c *Config) GetWriteMysql() *xorm.MysqlCfg {
	return &c.WriteMysql
}

//new config
func NewConfig(port int64, path string) *Config {

	//config
	var conf = &Config{
		ReadMysql:  xorm.MysqlCfg{},
		WriteMysql: xorm.MysqlCfg{},
		NatsCfg:    nats.NatsCfg{},
	}

	//read config
	err := ini.NewConfig().ReadConfig(path, conf)
	if err != nil {
		panic(err)
	}

	//init port
	if port > 0 {
		conf.Port = port
	}

	//get env port
	if p := os.Getenv("PORT"); p != "" {
		port, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			conf.Port = port
		}
	}

	if len(conf.WriteMysql.DbIP) <= 0 {
		conf.WriteMysql = conf.ReadMysql
	}

	fmt.Println("************************************************")
	fmt.Println("*                                              *")
	fmt.Println("*             	   Cfg  Init                    *")
	fmt.Println("*                                              *")
	fmt.Println("************************************************")
	fmt.Println("### Port:              ", conf.Port)
	fmt.Println("### ReadMysql:         ", conf.ReadMysql)
	fmt.Println("### WriteMysql:        ", conf.WriteMysql)
	fmt.Println("### NsqCfg:            ", conf.NsqCfg)
	fmt.Println("### NatsCfg:           ", conf.NatsCfg)
	fmt.Println("### RedisCfg:          ", conf.RedisCfg)
	return conf
}
