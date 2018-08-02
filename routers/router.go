package routers

import (
	"msgnotification/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/msgnotification",
		//404页面
		beego.NSRouter("/errorpage", &controllers.FourZeroController{}, "GET:Get"),

		beego.NSRouter("/postmessage", &controllers.MsgnotificationController{}, "POST:AddMessage"),
		beego.NSRouter("/getmessage", &controllers.MsgnotificationController{}, "POST:GetMessage"),
		beego.NSRouter("/getuser", &controllers.UserController{}, "POST:UserName"),
		beego.NSRouter("/setparam", &controllers.ParamController{}, "POST:SetParam"),
		beego.NSRouter("/sendmessage", &controllers.ViewController{}, "GET:SendMessage"),

	)
	beego.AddNamespace(ns)

}
