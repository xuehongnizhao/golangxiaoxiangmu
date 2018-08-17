package config

import (
"common/base"
"encoding/json"
"errors"
"fmt"
"io/ioutil"
"net/url"
"os"
"os/exec"
"path"
registryv1 "plugins/registry/v1"
"strings"
"time"

"github.com/astaxie/beego"
"github.com/astaxie/beego/httplib"
)

var (
	DefaultConfig *RunningConfig
	configAddress = ""
)

type RunningConfig struct {
	BeeConfig *beego.Config
	Base      baseConfig
	Log       logConfig       `json:"log"`
	FileLog   fileLogConfig   `json:"log_file"`
	MsgServer   msgServerConfig   `json:"msg_server"`
	SyslogLog syslogLogConfig `json:"log_syslog"`
	Registry  registryConfig  `json:"registry"`
	Server    serverConfig    `json:"server"`
	DataBase  dataBaseConfig  `json:"database_mysql"`
	Page      pageConfig      `json:"page"`
	Nsq       nsqConfig       `json:"nsq"`
	Template  templateConfig  `json:"template"`
}

type baseConfig struct {
	RootDir string
}
type msgServerConfig struct {
	Ip string
	Port string
}
type logConfig struct {
	Level    string `json:"level"`
	Type     string `json:"type"`
	LogLevel int
	LogType  []string //file,syslog
}
type fileLogConfig struct {
	LogDir  string `json:"dir"`
	LogFile string `json:"file"`
}

type registryConfig struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Mode    string `json:"mode"`
	Address string `json:"address"`
	TTL     int    `json:"ttl"`
}

type serverConfig struct {
	Ip   string
	Port int
}

type dataBaseConfig struct {
	Debug    bool
	Server   string
	Port     int
	Name     string
	User     string
	Password string
	Charset  string
	MaxCon   int `json:"maxcon"`
	MaxIdle  int `json:"maxidle"`
}
type pageConfig struct {
	Limit    int
	MaxLimit int
}

type syslogLogConfig struct {
}

type nsqConfig struct {
	Lookup      string `json:"lookup"`
	Lookups     []string
	Nsqds       []string
	Nsqd        string `json:"nsqds"`
	Topic       string `json:"topic"`
	ChannelName string `json:"channal"`
}
type templateConfig struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

func init() {
	DefaultConfig = new(RunningConfig)
	DefaultConfig.Base.RootDir, _ = exec.LookPath(os.Args[0])
	DefaultConfig.FileLog.LogDir = "/var/log/msgnotification"
	DefaultConfig.FileLog.LogFile = "msgnotification.log"
	DefaultConfig.Log.LogLevel = beego.LevelWarning
	DefaultConfig.Log.LogType = strings.Split("file", ",")
	DefaultConfig.Server.Ip = "127.0.0.1"
	DefaultConfig.Server.Port = 8080

}

//加载配置信息
func LoadConfig(cmd *Cmd) {
	if cmd.Config != "" {
		fmt.Printf("配置加载模式为：[ 本地文件 ]，文件名称为：[ %s ] \n", cmd.Config)
		LoadConfigFromFile(cmd.Config)
	} else {
		reg := registryv1.NewClient(cmd.Registry)
		regSvc, err := reg.GetService("gconfig")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("获取配置服务，服务地址为：[ %s ] \n", regSvc[0].Address)
		configAddress = fmt.Sprintf("http://%s/gconfig/v1/config", regSvc[0].Address)

		fmt.Printf("配置加载模式为：[ 远程服务 ]，服务地址为：[ %s ] \n", configAddress)
		result, err := queryConfig(cmd.ServiceName, cmd.Mode)
		if err != nil {
			fmt.Printf("读取配置信息失败，错误原因为：[ %s ] \n", err.Error())
			os.Exit(1)
		}

		strConf := result.Result["config"].(string)

		fmt.Println(fmt.Sprintf("配置信息为：[ %s ]", strConf))
		err = json.Unmarshal([]byte(strConf), DefaultConfig)
		if err != nil {
			fmt.Printf("加载服务配置信息失败，错误原因为：[ %s ] \n", err.Error())
			os.Exit(3)
		}
		//日志配置转换
		DefaultConfig.Log.LogLevel = convertLevel(DefaultConfig.Log.Level)
		DefaultConfig.Log.LogType = strings.Split(DefaultConfig.Log.Type, ",")
		//nsq配置转换
		//		DefaultConfig.Indicator.Lookups = strings.Split(DefaultConfig.Indicator.Lookup, ",")
		//		DefaultConfig.Indicator.Nsqds = strings.Split(DefaultConfig.Indicator.Nsqd, ",")

		//运行配置信息加载
		tmpName := path.Join(os.TempDir(), "beego", fmt.Sprintf("%d", time.Now().Nanosecond()))
		os.MkdirAll(path.Dir(tmpName), os.ModePerm)
		defer os.Remove(tmpName)
		if err := ioutil.WriteFile(tmpName, []byte(strConf), 0655); err != nil {
			fmt.Println(fmt.Sprintf("配置信息写入临时文件失败，错误内容为：[ %s ]", err.Error()))
			os.Exit(4)
		}

		err = beego.LoadAppConfig("json", tmpName)
		if err != nil {
			fmt.Println(fmt.Sprintf("加载配置信息失败，错误内容为：[ %s ]", err.Error()))
			os.Exit(5)
		}
	}
	beego.BConfig.WebConfig.TemplateLeft = DefaultConfig.Template.Left
	beego.BConfig.WebConfig.TemplateRight = DefaultConfig.Template.Right
	DefaultConfig.BeeConfig = beego.BConfig

	//打印当前配置信息
	printConfig()

}

//获取服务运行配置信息
func queryConfig(sname, mode string) (*base.ApiResult, error) {

	link, err := url.ParseRequestURI(configAddress)
	if err != nil {
		return nil, err
	}
	values := link.Query()
	values.Add("name", sname)
	values.Add("mode", mode)

	link.RawQuery = values.Encode()
	apiURL := link.String()
	beego.Debug("get config url is : ", apiURL)
	req := httplib.Get(apiURL)
	result := new(base.ApiResult)
	err = req.ToJSON(result)
	if err != nil {
		return nil, err
	}
	if result.Success {
		return result, nil
	} else {
		return nil, errors.New(result.Error.Msg)
	}
}

//加载配置文件
// ****  仅限在开发环境没有gconfig服务时，通过传递-c参数使用   *****
func LoadConfigFromFile(file string) {

	err := beego.LoadAppConfig("ini", file)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	DefaultConfig.Server.Ip = beego.AppConfig.DefaultString("server::ip", "127.0.0.1")
	DefaultConfig.Server.Port = beego.AppConfig.DefaultInt("server::port", 8080)
	DefaultConfig.Log.LogLevel = convertLevel(beego.AppConfig.DefaultString("log::level", "warn"))

	DefaultConfig.Log.LogType = strings.Split(beego.AppConfig.DefaultString("log::type", "file"), ",")

	DefaultConfig.Base.RootDir, _ = exec.LookPath(os.Args[0])

	DefaultConfig.DataBase.Debug = beego.AppConfig.DefaultBool("database_mysql::debug", false)
	DefaultConfig.DataBase.Server = beego.AppConfig.DefaultString("database_mysql::server", "192.168.10.137")
	DefaultConfig.DataBase.Port = beego.AppConfig.DefaultInt("database_mysql::port", 3306)
	DefaultConfig.DataBase.Name = beego.AppConfig.DefaultString("database_mysql::name", "onecard")
	DefaultConfig.DataBase.User = beego.AppConfig.DefaultString("database_mysql::user", "openfire")
	DefaultConfig.DataBase.Password = beego.AppConfig.DefaultString("database_mysql::password", "123456")
	DefaultConfig.DataBase.Charset = beego.AppConfig.DefaultString("database_mysql::charset", "utf8")
	DefaultConfig.DataBase.MaxCon = beego.AppConfig.DefaultInt("database_mysql::max_conn", 20)
	DefaultConfig.DataBase.MaxIdle = beego.AppConfig.DefaultInt("database_mysql::max_idle", 5)

	DefaultConfig.Page.Limit = beego.AppConfig.DefaultInt("page::limit", 20)
	DefaultConfig.Page.MaxLimit = beego.AppConfig.DefaultInt("page::max_limit", 1000)

	DefaultConfig.Registry.Enabled = beego.AppConfig.DefaultBool("registry::enabled", true)
	DefaultConfig.Registry.Address = beego.AppConfig.DefaultString("registry::address", "192.168.10.127")
	DefaultConfig.Registry.Mode = beego.AppConfig.DefaultString("registry::mode", "dev")
	DefaultConfig.Registry.Name = beego.AppConfig.DefaultString("registry::name", "canteen")
	DefaultConfig.Registry.TTL = beego.AppConfig.DefaultInt("registry::ttl", 60)

	DefaultConfig.Template.Left = beego.AppConfig.DefaultString("template::left", "<<<")
	DefaultConfig.Template.Right = beego.AppConfig.DefaultString("template:right", ">>>")
	DefaultConfig.MsgServer.Ip = beego.AppConfig.DefaultString("msg_server::ip", "10.166.1.29")
	DefaultConfig.MsgServer.Port = beego.AppConfig.DefaultString("msg_server::port", "10000")

}

func convertLevel(level string) int {
	switch level {
	case "debug":
		return beego.LevelDebug
	case "info":
		return beego.LevelInformational
	case "warn":
		return beego.LevelWarning
	case "error":
		return beego.LevelError
	}
	return beego.LevelWarning
}

//获取当前工作目录
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
		return "/"
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func printConfig() {
	// 打印输出配置信息
	fmt.Println("************** DefaultConfig BeeConfig  ******************")
	fmt.Println(fmt.Sprintf("%+v", DefaultConfig.BeeConfig))
	fmt.Println("************** DefaultConfig BeeConfig over **************")
	fmt.Println("")
	fmt.Println("************** DefaultConfig *******************")
	fmt.Println(fmt.Sprintf("%+v", DefaultConfig))
	fmt.Println("************** DefaultConfig over **************")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

}
