package main

import (
	"fmt"
	"msgnotification/config"
	"msgnotification/log"
	"msgnotification/models"
	"msgnotification/plugin/registry"
	_ "msgnotification/routers"

	"github.com/astaxie/beego"
)

func main() {
	cmd := new(config.Cmd)
	//解析 传入参数
	cmd.Parse()

	//加载配置信息
	config.LoadConfig(cmd)

	//启动日志服务
	log.InitLog()

	//注册数据库服务
	models.InitDB()
	//增加静态文件路径
	beego.SetStaticPath("/msgnotification/static", "static")
	//初始化rigistry服务
	registry.InitRegistry(cmd)

	beego.Run(fmt.Sprintf("%s:%d", config.DefaultConfig.BeeConfig.Listen.HTTPAddr, config.DefaultConfig.BeeConfig.Listen.HTTPPort))
}
