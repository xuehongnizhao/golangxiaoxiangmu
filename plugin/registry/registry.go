package registry

import (
	"msgnotification/config"
	registry "plugins/registry/v1"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	//最小同步周期，单位秒
	MIN_PERIOD = 5
	//注册服务周期缩短偏移量，单位秒
	OFFSET = 2
)

var (
	client *registry.Client
	mode   string
)

func InitRegistry(cmd *config.Cmd) {
	client = registry.NewClient(cmd.Registry)
	mode = config.DefaultConfig.Registry.Mode
	if config.DefaultConfig.Registry.Enabled {
		service := registry.Service{}
		service.Name = config.DefaultConfig.Registry.Name
		service.Mode = config.DefaultConfig.Registry.Mode
		service.Address = config.DefaultConfig.Registry.Address
		service.TTL = config.DefaultConfig.Registry.TTL
		err := RegistService(service)
		if err != nil {
			beego.Error(err)
		}
	}

}

func formatEndpoint(address string) string {

	if !strings.HasPrefix(address, "http") {
		return "http://" + address
	}
	return address
}

//功能说明:registry注册服务地址和端口
//创建人：高勇
//创建时间：2016-08-09 10:58:11
func RegistService(svc registry.Service) error {

	if svc.TTL < MIN_PERIOD {
		svc.TTL = MIN_PERIOD
	}
	_, err := client.RegistService(svc)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Second * time.Duration(svc.TTL-OFFSET))
	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := client.RegistService(svc)
				if err != nil {
					beego.Error(err)
				}
			}

		}
	}()

	return nil
}

func GetCMDBCenterService() (string, error) {
	regSvc, err := client.GetService("cmdb", mode)
	if err != nil {
		return "", err
	}
	return regSvc[0].Address, nil
}

//获取canteen服务地址
func GetMonitorCenterService() (val string, err error) {
	//	regSvc, err := client.GetService("canteen", mode)
	//	if err != nil {
	//		return "", err
	//	}
	return "127.0.0.1:8080", nil
}

//获取attendance服务注册地址
func GetAttendanceService() (val string, err error) {
	regSvc, err := client.GetService("attendance", mode)
	if err != nil {
		return "", err
	}
	beego.Debug("==============", regSvc[0].Address)
	return regSvc[0].Address, nil
}

//获取stopenfire服务地址
func GetStopenfireService() (val string, err error) {
	regSvc, err := client.GetService("stopenfire", mode)
	if err != nil {
		return "", err
	}
	beego.Debug("==============", regSvc[0].Address)
	return regSvc[0].Address, nil
}
