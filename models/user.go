package models

import (
"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int64
	Name   string
	Rtxaccount   string
}

func init() {
	orm.RegisterModel(new(User))
}


//查人
func UserName() ([]*User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	
	user := make([]*User, 0)
 	_,err := qs.All(&user)

	return user, err
}
