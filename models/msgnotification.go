package models

import (
"common/base"
"fmt"
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
}

func init() {
	orm.RegisterModel(new(Msgnotification))
}

type QueryMsgnotification struct {
	BaseOption *base.QueryOptions
	Name       string
	Date      string
}


//查询发送记录
func QueryMessage(opt *QueryMsgnotification) ([]*Msgnotification, int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Msgnotification))
	cond := new(orm.Condition)

	if opt.Name != "" {
		cond = cond.And("name__icontains", opt.Name)
	}
	if opt.Date != "" {
		cond = cond.And("date", opt.Date)
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

//添加发送记录
func PostMessage(dt *Msgnotification) (*Msgnotification, error) {
	o := orm.NewOrm()
	o.Begin()
	err := dt.Valited()
	id, err := o.Insert(dt)
	if err != nil {
		o.Rollback()
		return nil, err
	}
	o.Commit()
	dt.Id = id
	return dt, err
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

