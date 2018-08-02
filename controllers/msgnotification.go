package controllers

import (
	"common/ajax"
	"common/base"
	"fmt"
	"msgnotification/models"

	"github.com/astaxie/beego"
)

type MsgnotificationController struct {
	beego.Controller
}

//查询信息列表
func (this *MsgnotificationController) GetMessage() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar

	opt := new(models.QueryMsgnotification)
	bp := new(base.QueryOptions)

	opt.BaseOption = bp
	var err error
	limit, err := this.GetInt("limit", 0)
	if err != nil {
		ar.SetError("limit获取发生异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	page, err := this.GetInt("page", 1)
	if err != nil {
		ar.SetError("page获取发生异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	opt.BaseOption.Limit = limit
	fmt.Println("-----%d",limit)
	opt.BaseOption.Offset = limit * (page - 1)
	opt.Name = this.GetString("name", "")
	opt.Date = this.GetString("date","")
	message, num, err := models.QueryMessage(opt)
	if err != nil {
		ar.SetError("获取部门信息发生异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}

	ar.Total = num
	ar.Data = message
	ar.Success = true
	this.ServeJSON()
}

//添加信息
func (this *MsgnotificationController) AddMessage() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar

	dt := new(models.Msgnotification)
	name := this.GetString("name", "")
	if name=="" {
		ar.SetError(fmt.Sprintf("对不起你没有权限这么做"))
		this.ServeJSON()
		return
	}
	content :=this.GetString("content", "")
	ending := this.GetString("ending", "")
	tel := this.GetString("tel", "")
	dt.Name = name
	dt.Ending=ending
	dt.Tel=tel
	dt.Content=content
	dt.Date = base.GetCurrentData()
	dtnew, err := models.PostMessage(dt)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加记录发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Data = dtnew
	ar.Success = true
	this.ServeJSON()
}
