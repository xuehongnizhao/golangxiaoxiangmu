package controllers

import (
"common/ajax"
"common/base"
"fmt"
"msgnotification/models"
"net/url"
"strings"
"time"
"github.com/astaxie/beego"
"github.com/astaxie/beego/httplib"
"msgnotification/config"
)

type MsgnotificationController struct {
	beego.Controller
}
type sendStatus struct {
	SStatus int
	FStatus int
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
	opt.Date = this.GetString("date", "")
	opt.Ftime = this.GetString("ftime", "")
	opt.Ltime = this.GetString("ltime", "")
	opt.Tel = this.GetString("tel", "")
	message, num, err := models.QueryMessage(opt)
	if err != nil {
		ar.SetError("获取信息发生异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	for _, v := range message {
		telarr, err := models.GetTel(v.Id)
		if err != nil {
			ar.SetError("获取电话号异常")
			beego.Error(err)
			this.ServeJSON()
			return
		}
		v.Telnumber = telarr
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
	if name == "" {
		ar.SetError(fmt.Sprintf("对不起你没有权限这么做"))
		this.ServeJSON()
		return
	}
	content := this.GetString("content", "")
	ending := this.GetString("ending", "")
	tel := this.GetString("tel", "")
	msg.Name = name
	msg.Ending = ending
	msg.Tel = tel
	msg.Status = 0
	msg.Content = content
	msg.Date = base.GetCurrentData()
	pid, err := models.PostMessage(msg)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加记录发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	numberarr := strings.Split(tel, ";")
	telnumberarr := make([]*models.Telnumber, 0)
	sendSuccess := true
	statuscode := "200 OK"
	for i := 0; i < len(numberarr); i++ {
		if numberarr[i] != "" {
			telnumbermodel := new(models.Telnumber)
			telnumbermodel.Pid = pid
			telnumbermodel.Tel = numberarr[i]
			statuscode = sendMsgWithService(numberarr[i], content+ending)
			if statuscode != "200 OK" {
				telnumbermodel.Status = -1
				sendSuccess = false
			} else {
				telnumbermodel.Status = 1
			}

			telnumberarr = append(telnumberarr, telnumbermodel)
		}
	}
	err = models.AddMsgTel(telnumberarr)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加记录发生异常，错误原因为：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	if !sendSuccess {
		msg.Status = -1
		err := models.UpdateMsgNotification(msg)
		if err != nil {
			ar.SetError(fmt.Sprintf("更新状态异常，错误原因为：[%s]", statuscode))
			ar.Success = false
			this.ServeJSON()
			return
		}
		ar.SetError(fmt.Sprintf("发送短信异常，错误原因为：[%s]", statuscode))
		ar.Success = false
		this.ServeJSON()
		return
	}
	msg.Status = 1
	err = models.UpdateMsgNotification(msg)
	if err != nil {
		ar.SetError(fmt.Sprintf("更新状态异常，错误原因为：[%s]", statuscode))
		ar.Success = false
		this.ServeJSON()
		return
	}
	ar.SetError(fmt.Sprintf("发送成功"))
	ar.Success = true
	this.ServeJSON()

}
func (this *MsgnotificationController) ResendMessage() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar

	id, err := this.GetInt64("id", 0)
	if err != nil {
		ar.SetError("电话号码获取异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}
	tel := new(models.Telnumber)
	tel, err = models.GetTelWithId(id)
	if err != nil {
		ar.SetError("查询电话号码异常")
		beego.Error(err)
		this.ServeJSON()
		return
	}

	statuscode := sendMsgWithService(tel.Tel, tel.Content)
	if statuscode != "200 OK" {
		ar.SetError(fmt.Sprintf("发送短信异常，错误原因为：[%s]", statuscode))
		this.ServeJSON()
		return
	}
	tel.Status = 1
	fStatus, err := models.UpdateTelStatus(tel)
	statusNew := new(sendStatus)
	statusNew.FStatus = fStatus
	statusNew.SStatus = 1
	ar.Data = statusNew
	ar.SetError(fmt.Sprintf("发送成功"))
	ar.Success = true
	this.ServeJSON()
}
func sendMsgWithService(sendNumber string, sendContent string) string {
	serverurl := "htttp://"+config.DefaultConfig.MsgServer.Ip+":"+config.DefaultConfig.MsgServer.Port+"/SMSService/SMSrestful/sendMessage"
	beego.Debug(serverurl)
	req := httplib.Post(serverurl)
	req.Header("Content-Type", "Application/json")
	req.Param("sendNum", sendNumber)
	contentP := url.QueryEscape(sendContent)
	req.Param("sendContent", contentP)
	req.Param("sourceSystemName", "jxn")
	resp, _ := req.SetTimeout(5 * time.Second, 5 * time.Second).Response()
	if resp==nil {
		return "短信服务发生错误"
	}
	return resp.Status
}
