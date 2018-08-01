package log

import (
	"encoding/json"
	"fmt"
	"msgnotification/config"
	"os"
	"path"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/logs"
)

type multiFileConfig struct {
	Filename string   `json:"filename"`
	Separate []string `json:"separate"`
}

func InitLog() {

	// 设置日志输出级别
	beego.SetLevel(config.DefaultConfig.Log.LogLevel)
	//设置输出文件名和行号
	beego.SetLogFuncCall(true)
	//设置显示层级
	//logs.SetLogFuncCallDepth(4)

	// 重置所有已注册的日志输出,防止二次注册时的错误
	beego.BeeLogger.Reset()

	for _, lt := range config.DefaultConfig.Log.LogType {
		switch lt {
		case "console":

			err := beego.SetLogger(logs.AdapterConsole, "")
			if err != nil {
				beego.Error(err)
				os.Exit(3)
			}
		case "syslog":

		case "file":
			logdir := config.DefaultConfig.FileLog.LogDir

			err := os.MkdirAll(logdir, 0755)
			if err != nil {
				fmt.Println(fmt.Sprintf("创建日志目录错误,目录为:[ %s ],错误内容为： %s", logdir, err.Error()))
				os.Exit(3)
			}
			config.DefaultConfig.FileLog.LogFile = path.Join(logdir, config.DefaultConfig.FileLog.LogFile)

			c := multiFileConfig{}
			c.Filename = config.DefaultConfig.FileLog.LogFile
			//fmt.Println(c.Filename)
			c.Separate = []string{"error", "warn", "info", "debug"}
			configBys, _ := json.Marshal(c)
			err = beego.SetLogger(logs.AdapterFile, string(configBys))
			if err != nil {
				beego.Error(err)
				os.Exit(3)
			}
		}

	}

	logtypes := strings.Join(config.DefaultConfig.Log.LogType, ",")
	beego.Debug(fmt.Sprintf("日志服务初始化成功,启动日志模块包括： %s, 日志级别为 %d", logtypes, config.DefaultConfig.Log.LogLevel))

}
