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


//通过rtxaccount查人
func UserName(rtxaccount string) ([]*User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	
	user := make([]*User, 0)
 	_,err := qs.Filter("rtxaccount",rtxaccount).All(&user)

	return user, err
}
