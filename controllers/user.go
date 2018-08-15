package controllers

import (
"common/ajax"
"msgnotification/models"

"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
type ReturnData struct{
	Name string
	NameArr []*models.User
}
//查询人员
func (this *UserController) UserName() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	user, err := models.UserName()
	if err!=nil {
		ar.SetError("错误请稍后重试或联系管理员")
		this.ServeJSON()
		return
	}
	returnData:=new(ReturnData)
	returnData.NameArr = user
	isTrue:=false
	for i := 0; i < len(user); i++ {
		if user[i].Rtxaccount==this.GetString("rtxaccount","") {
			isTrue=true
			returnData.Name = user[i].Name

			break
		}
	}
	if isTrue {
		ar.Data = returnData
		ar.Success = true
		this.ServeJSON()
		return
	}else{
		ar.SetError("对不起权限受限 请联系管理员")
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}
