package models

import (
	"common/base"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Msgnotification struct {
	Id        int64
	Name      string
	Date      string
	Content   string
	Tel       string
	Ending    string
	Status    int64
	Telnumber []*TelnumberOption `orm:"-"`
}

type Telnumber struct {
	Id      int64
	Pid     int64
	Tel     string
	Status  int64
}

type TelnumberOption struct {
	Id      int64
	Pid     int64
	Tel     string
	Status  int64
	Content string 
	Name string 
}



func init() {
	orm.RegisterModel(new(Msgnotification), new(Telnumber))
}

type QueryMsgnotification struct {
	BaseOption *base.QueryOptions
	Name       string
	Ftime      string
	Ltime      string
	Tel        string
	Date       string
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
		cond = cond.And("name", opt.Name)
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
func GetTel(opt int64) ([]*TelnumberOption, error) {
	o := orm.NewOrm()
	sql := "select a.id,a.pid,a.status,a.tel,b.name from telnumber a left JOIN contacts b on a.tel=b.tel where pid = "+strconv.FormatInt(opt, 10) +";"
	msg := make([]*TelnumberOption,0)

	_,err := o.Raw(sql).QueryRows(&msg)
	return msg, err
}

func GetTelWithId(id int64) (*TelnumberOption, error) {
	o := orm.NewOrm()
	sql := "SELECT b.`id`,b.pid, a.`content` ,b.`tel` FROM msgnotification a, telnumber b WHERE a.`id`=b.`pid` AND b.`id`=" + strconv.FormatInt(id, 10) + ";"
	msg := new(TelnumberOption)
	err := o.Raw(sql).QueryRow(&msg)
	return msg, err
}

//添加发送记录
func PostMessage(msg *Msgnotification) (int64, error) {
	o := orm.NewOrm()
	o.Begin()
	err := msg.Valited()
	id, err := o.Insert(msg)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return -1, err
	}
	o.Commit()

	return id, err
}
func AddMsgTel(tel []*Telnumber) error {
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
func UpdateTelStatus(tel *Telnumber) (int, error) {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(tel, "status")
	if err != nil {
		o.Rollback()
		return -1, err
	}
	telM := make([]*Telnumber, 0)
	sql := "SELECT * FROM telnumber WHERE pid =" + strconv.FormatInt(tel.Pid, 10) + " AND STATUS=-1;"
	count, err := o.Raw(sql).QueryRows(&telM)
	if count == 0 {
		msg := new(Msgnotification)
		msg.Id = tel.Pid
		msg.Status = 1
		_, err = o.Update(msg, "status")
		if err != nil {
			o.Rollback()
			return -1, err
		}
		o.Commit()
		return 1, err
	}

	o.Commit()
	return -1, err

}
func UpdateMsgNotification(msg *Msgnotification) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(msg, "status")
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return err

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
