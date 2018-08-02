package controllers

import (
"common/ajax"
"fmt"
"msgnotification/models"

"github.com/astaxie/beego"
)

type ParamController struct {
	beego.Controller
}

func (this *ParamController) SetParam() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	key := this.GetString("key", "")
	value := this.GetString("value", "")
	
	err := models.SetParam(key,value)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加公版结尾发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}
