package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Cmd struct {
	ServiceName string `json:"servicename"`
	Mode        string `json:"mode"`
	Config      string `json:"config"`
	Registry    string `json:"registry"`
}

//解析传入参数
func (c *Cmd) Parse() {

	//读取传递参数
	c.parseArgs()

	//读取环境变量
	c.parseEnv()

}

func (c *Cmd) parseArgs() {
	//读取传入参数，如果没有传递则默认读取环境变量
	flag.StringVar(&c.ServiceName, "s", c.ServiceName, "servicename in gconfig")
	flag.StringVar(&c.Mode, "m", c.Mode, "runmode in gconfig")
	flag.StringVar(&c.Config, "c", c.Config, "file name of config")
	flag.StringVar(&c.Registry, "r", c.Registry, "registry service address")

	flag.Parse()
	log.Println("输入参数读取后：")
	log.Println(c.String())
}

func (c *Cmd) parseEnv() {

	log.Println("RGS_SERVICE_NAME : ", os.Getenv("RGS_SERVICE_NAME"))
	log.Println("RGS_MODE : ", os.Getenv("RGS_MODE"))
	log.Println("RGS_CONFIG : ", os.Getenv("RGS_CONFIG"))
	log.Println("RGS_REGISTRY : ", os.Getenv("RGS_REGISTRY"))

	c.ServiceName = checkEmpty(os.Getenv("RGS_SERVICE_NAME"), c.ServiceName)
	c.Mode = checkEmpty(os.Getenv("RGS_MODE"), c.Mode)
	c.Config = checkEmpty(os.Getenv("RGS_CONFIG"), c.Config)
	c.Registry = checkEmpty(os.Getenv("RGS_REGISTRY"), c.Registry)
	log.Println("环境变量读取后：")
	log.Println(c.String())

}

func (c *Cmd) String() string {
	bys, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(bys)
}

func checkEmpty(val, defVal string) string {
	if val == "" {
		return defVal
	}
	return val
}
