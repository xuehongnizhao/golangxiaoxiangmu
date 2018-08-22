package models

import (
"github.com/astaxie/beego"
"github.com/astaxie/beego/orm"
)

type Contacts struct {
	Id        int64
	Name      string
	Cdate      string
	Tel       string
	Status    int64
}

func init() {
	orm.RegisterModel(new(Contacts))
}

//查询联系人
func QueryContacts(status int64) ([]*Contacts, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Contacts))
	if status==1 {
		cond := new(orm.Condition)
		cond = cond.And("status", 1)
		qs = qs.SetCond(cond)
	}
	contacts := make([]*Contacts, 0)
	_, err := qs.OrderBy("-status","id","name").All(&contacts)
	return contacts, err
}

//添加联系人
func AddContact(contact *Contacts)  error{
	o := orm.NewOrm()
	o.Begin()
	_,err := o.Insert(contact)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}
func UpdateContact(contact *Contacts)  error {
	beego.Debug(contact)
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(contact)
	if err != nil {
		o.Rollback()
		return  err
	}
	o.Commit()
	return  err

}
func IsTelExis(tel string) (int64 ,error){

	o := orm.NewOrm()
	contacts := make([]*Contacts, 0)
	num,err := o.QueryTable(new(Contacts)).Filter("Tel", tel).All(&contacts)
	if num>0 {
		return contacts[0].Id,err
	}
	return -1,err
}