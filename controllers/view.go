/*-------页面显示路由-------*/

package controllers

import (
	"msgnotification/models"
	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) SendMessage() {
	this.TplName = "sendMessage.html"
	rtxaccount := this.GetString("rtxaccount","")
	ending ,err:= models.QueryParam("ending")
	if err!=nil {
		
	}
	this.Data["rtxaccount"] = rtxaccount
	this.Data["ending"] = ending
}
