package routers

import (
	"msgnotification/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/xt/msgnotification",
		//404页面
		beego.NSRouter("/errorpage", &controllers.FourZeroController{}, "GET:Get"),
	)
	beego.AddNamespace(ns)

}
