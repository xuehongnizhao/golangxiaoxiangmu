package controllers

import (
"common/ajax"
"common/base"
"fmt"
"msgnotification/models"
"strings"
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
	opt.BaseOption.Offset = limit * (page - 1)
	opt.Name = this.GetString("name", "")
	opt.Date = this.GetString("date","")
	opt.Ftime = this.GetString("ftime","")
	opt.Ltime = this.GetString("ltime","")
	opt.Tel = this.GetString("tel","")
	fmt.Println(opt.Name,opt.Ftime)
	message, num, err := models.QueryMessage(opt)
	if err != nil {
		ar.SetError("获取信息发生异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	for _,v := range(message){
		telarr ,err := models.GetTel(v.Id)
		if err != nil {
			ar.SetError("获取电话号异常")
			beego.Error(err)
			this.ServeJSON()
			return
		}
		v.Telnumber=telarr
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

	msg := new(models.Msgnotification)
	name := this.GetString("name", "")
	if name=="" {
		ar.SetError(fmt.Sprintf("对不起你没有权限这么做"))
		this.ServeJSON()
		return
	}
	content :=this.GetString("content", "")
	ending := this.GetString("ending", "")
	tel := this.GetString("tel", "")
	msg.Name = name
	msg.Ending=ending
	msg.Tel=tel
	msg.Content=content
	msg.Date = base.GetCurrentData()
	pid,err := models.PostMessage(msg)
	numberarr := strings.Split(tel,";")
	telnumberarr := make([]*models.Telnumber, 0)
	for i := 0; i < len(numberarr); i++ {
		if numberarr[i]!="" {
			telnumbermodel:=new(models.Telnumber)
			telnumbermodel.Pid=pid
			telnumbermodel.Tel=numberarr[i]
			telnumberarr=append(telnumberarr,telnumbermodel)
		}
	}
	
	err = models.AddMsgTel(telnumberarr)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加记录发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}
