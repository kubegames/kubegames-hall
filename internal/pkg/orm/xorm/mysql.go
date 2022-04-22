package xorm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"xorm.io/core"
	basexorm "xorm.io/xorm"
)

type (
	//cfg
	Cfg interface {
		//get read mysql
		GetReadMysql() *MysqlCfg

		//GetWriteMysql get write mysql
		GetWriteMysql() *MysqlCfg
	}

	//mysql config
	MysqlCfg struct {
		//db ip
		DbIP string `ini:"DbIp"`
		//db pwd
		DbPWd string `ini:"DbPwd"`
		//db user
		DbUser string `ini:"DbUser"`
		//database
		Database string `ini:"Database"`
		//SetMaxIdleConns set the max idle connections on pool, default is 2
		MaxIdleConns int `ini:"MaxIdleConns"`
		//SetMaxOpenConns is only available for go 1.2+
		MaxOpenConns int `ini:"MaxOpenConns"`
		//SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		ConnMaxLifetime int `ini:"ConnMaxLifetime"`
		//log switch
		OpenLog bool `ini:"OpenLog"`
	}
)

//mysql
type Mysql struct {
	read  *basexorm.Engine
	write *basexorm.Engine
}

//NewMysql new mysql
func NewMysql(cfg Cfg) *Mysql {
	mysql := new(Mysql)

	//init read mysql
	readurl := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		cfg.GetReadMysql().DbUser,
		cfg.GetReadMysql().DbPWd,
		cfg.GetReadMysql().DbIP,
		cfg.GetReadMysql().Database)
	log.Infof("read mysql %s", readurl)
	read, err := basexorm.NewEngine("mysql", readurl)
	if err != nil {
		panic("connect to mysql error:" + err.Error())
	}
	if cfg.GetReadMysql().MaxIdleConns > 0 {
		read.SetMaxIdleConns(cfg.GetReadMysql().MaxIdleConns)
	}
	if cfg.GetReadMysql().MaxOpenConns > 0 {
		read.SetMaxOpenConns(cfg.GetReadMysql().MaxOpenConns)
	}
	if cfg.GetReadMysql().ConnMaxLifetime > 0 {
		read.SetConnMaxLifetime(time.Duration(cfg.GetReadMysql().ConnMaxLifetime) * time.Second)
	}
	read.SetLogger(new(Logger))
	read.ShowSQL(cfg.GetReadMysql().OpenLog)
	read.SetTableMapper(core.GonicMapper{})
	read.SetColumnMapper(core.GonicMapper{})
	mysql.read = read

	//init write mysql
	writeurl := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		cfg.GetWriteMysql().DbUser,
		cfg.GetWriteMysql().DbPWd,
		cfg.GetWriteMysql().DbIP,
		cfg.GetWriteMysql().Database)
	log.Infof("write mysql %s", writeurl)
	write, err := basexorm.NewEngine("mysql", writeurl)
	if err != nil {
		panic("connect to mysql error:" + err.Error())
	}
	if cfg.GetWriteMysql().MaxIdleConns > 0 {
		write.SetMaxIdleConns(cfg.GetWriteMysql().MaxIdleConns)
	}
	if cfg.GetWriteMysql().MaxOpenConns > 0 {
		write.SetMaxOpenConns(cfg.GetWriteMysql().MaxOpenConns)
	}
	if cfg.GetWriteMysql().ConnMaxLifetime > 0 {
		write.SetConnMaxLifetime(time.Duration(cfg.GetWriteMysql().ConnMaxLifetime) * time.Second)
	}

	//init set
	write.SetLogger(new(Logger))
	write.ShowSQL(cfg.GetWriteMysql().OpenLog)
	write.SetTableMapper(core.GonicMapper{})
	write.SetColumnMapper(core.GonicMapper{})

	mysql.write = write
	return mysql
}

//return read Mysql engine
func (m *Mysql) GetReadEngine() *basexorm.Engine {
	return m.read
}

//return write Mysql engine
func (m *Mysql) GetWriteEngine() *basexorm.Engine {
	return m.write
}
