package routers

import (
	"msgnotification/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/xt/msgnotification",
		//404页面
		beego.NSRouter("/errorpage", &controllers.FourZeroController{}, "GET:Get"),
		//添加发送记录
		beego.NSRouter("/postmessage", &controllers.MsgnotificationController{}, "POST:AddMessage"),
		//获取发送记录
		beego.NSRouter("/getmessage", &controllers.MsgnotificationController{}, "POST:GetMessage"),
		//重新发送短信
		beego.NSRouter("/resendmessage", &controllers.MsgnotificationController{}, "POST:ResendMessage"),
		//获取用户姓名
		beego.NSRouter("/getuser", &controllers.UserController{}, "POST:UserName"),
		//设置短信结尾
		beego.NSRouter("/setparam", &controllers.ParamController{}, "POST:SetParam"),
		beego.NSRouter("/", &controllers.ViewController{}, "GET:SendMessage"),
	)
	beego.AddNamespace(ns)

}
