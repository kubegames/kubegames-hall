package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/pkg/config"
)

func main() {
	//new config
	config := config.NewConfig(8080, "./config.ini")

	//new mysql
	mysql := xorm.NewMysql(config)

	//new
	NewCasbin(mysql)
}

func NewCasbin(mysql *xorm.Mysql) {
	// a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/casbin")
	// if err != nil {
	// 	panic(err)
	// }

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(m, mysql.GetWriteEngine())
	if err != nil {
		panic(err)
	}

	//添加规则
	e.AddPolicy("admin", "/test", "GET")
}
