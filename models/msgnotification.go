package models

import (
"common/base"
"fmt"
"github.com/astaxie/beego"
"github.com/astaxie/beego/orm"
"github.com/astaxie/beego/validation"
)

type Msgnotification struct {
	Id     int64
	Name   string
	Date   string
	Content string
	Tel string
	Ending string
	Status int64
	Telnumber []*Telnumber `orm:"-"`
}

type Telnumber struct {
	Id     int64
	Pid   int64
	Tel   string
	Status int64
}

func init() {
	orm.RegisterModel(new(Msgnotification),new(Telnumber))
}

type QueryMsgnotification struct {
	BaseOption *base.QueryOptions
	Name       string
	Ftime 	string
	Ltime	string
	Tel 	string
	Date  	string
}


//查询发送记录
func QueryMessage(opt *QueryMsgnotification) ([]*Msgnotification, int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Msgnotification))
	cond := new(orm.Condition)

	if opt.Tel != "" {
		cond = cond.And("tel__icontains", opt.Tel)
	}
	if opt.Name != "" {
		cond = cond.And("name__icontains", opt.Name)
	}
	if opt.Ftime != "" {
		cond = cond.And("date__gte", opt.Ftime)
	}
	if opt.Ltime != "" {
		cond = cond.And("date__lte", opt.Ltime)
	}
	qs = qs.SetCond(cond)
	message := make([]*Msgnotification, 0)
	num, err := qs.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = qs.OrderBy("-date").Limit(opt.BaseOption.Limit).Offset(opt.BaseOption.Offset).All(&message)
	return message, int(num), err
}
func GetTel(opt int64) ([]*Telnumber, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Telnumber))
	cond := new(orm.Condition)
	cond = cond.And("pid", opt)
	qs = qs.SetCond(cond)
	telnumber := make([]*Telnumber, 0)
	_, err := qs.All(&telnumber)
	return telnumber, err
}

//添加发送记录
func PostMessage(msg *Msgnotification) ( int64, error) {
	o := orm.NewOrm()
	o.Begin()
	err := msg.Valited()
	id, err := o.Insert(msg)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return -1,err
	}
	o.Commit()

	return id,err
}
func AddMsgTel(tel []*Telnumber)  error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.InsertMulti(len(tel), tel)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}
	o.Commit()

	return nil
}


func (this *Msgnotification) Valited() error {
	valid := validation.Validation{}
	valid.Required(this.Name, "name")
	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

