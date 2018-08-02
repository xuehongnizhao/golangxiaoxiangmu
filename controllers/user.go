package controllers

import (
"common/ajax"
"msgnotification/models"

"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//查询人员
func (this *UserController) UserName() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	user, err := models.UserName(this.GetString("rtxaccount",""))

	if err != nil || len(user)<=0 {
		ar.SetError("权限受限，请联系管理员")
		this.ServeJSON()
		return
	}
	ar.Data = user
	ar.Success = true
	this.ServeJSON()
}
