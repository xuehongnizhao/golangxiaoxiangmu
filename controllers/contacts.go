package controllers

import (
"common/ajax"
"common/base"
"msgnotification/models"
"fmt"
"github.com/astaxie/beego"
)

type ContactsController struct {
	beego.Controller
}

//查询人员
func (this *ContactsController) GetContact() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	status,err:=this.GetInt64("status",0)
	if err != nil {
		ar.SetError("状态获取异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}

	contacts, err := models.QueryContacts(status)
	if err!=nil {
		ar.SetError("错误请稍后重试或联系管理员")
		this.ServeJSON()
		return
	}
	ar.Data=contacts
	ar.Success = true
	this.ServeJSON()
}
//新增人员
func (this *ContactsController) AddContact() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	contact := new(models.Contacts)
	name := this.GetString("name", "")
	tel := this.GetString("tel", "")
	isExist ,err:= models.IsTelExis(tel)
	if err!=nil {
		ar.SetError(fmt.Sprintf("添加人员发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	if isExist>0 {
		ar.SetError(fmt.Sprintf("电话号已经存在不能重复添加"))
		this.ServeJSON()
		return
	}
	contact.Name = name
	contact.Tel = tel
	contact.Status = 1
	contact.Cdate = base.GetCurrentData()
	err = models.AddContact(contact)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加人员发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}
func (this *ContactsController) UpdateContact() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	contact := new(models.Contacts)
	name := this.GetString("name", "")
	tel := this.GetString("tel", "")
	id,err:=this.GetInt64("id",0)
	if err != nil {
		ar.SetError("id获取异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	status,err:=this.GetInt64("status",0)
	if err != nil {
		ar.SetError("status获取异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	isExist ,err:= models.IsTelExis(tel)
	if err!=nil {
		ar.SetError(fmt.Sprintf("添加人员发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}

	if isExist>0&&isExist!=id {
		ar.SetError(fmt.Sprintf("电话号已经存在不能修改"))
		this.ServeJSON()
		return
	}
	contact.Id = id
	contact.Name = name
	contact.Tel = tel
	contact.Status = status
	contact.Cdate = base.GetCurrentData()
	err = models.UpdateContact(contact)
	if err != nil {
		ar.SetError(fmt.Sprintf("修改人员发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}
