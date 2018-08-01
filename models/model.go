package models

import (
	"fmt"
	"msgnotification/config"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	beego.Debug("正在初始化数据库连接信息...")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	dburl := "%s:%s@tcp(%s:%d)/%s?charset=%s"
	dburl = fmt.Sprintf(dburl,
		config.DefaultConfig.DataBase.User,
		config.DefaultConfig.DataBase.Password,
		config.DefaultConfig.DataBase.Server,
		config.DefaultConfig.DataBase.Port,
		config.DefaultConfig.DataBase.Name,
		config.DefaultConfig.DataBase.Charset)
	beego.Informational(fmt.Sprintf("数据库连接信息为： %s", dburl))
	err = orm.RegisterDataBase("default", "mysql", dburl, config.DefaultConfig.DataBase.MaxIdle, config.DefaultConfig.DataBase.MaxCon)
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	orm.Debug = config.DefaultConfig.DataBase.Debug
}
