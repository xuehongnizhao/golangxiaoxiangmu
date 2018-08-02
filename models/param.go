package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Param struct {
	Id    int64
	Key   string
	Value string
}

func init() {
	orm.RegisterModel(new(Param))
}

//设置系统参数（key=status时 1 普通状态 2 检查状态）
func SetParam(key, value string) error {
	o := orm.NewOrm()
	o.Begin()
	qs := o.QueryTable("param").Filter("key", key)
	if qs.Exist() {
		_, err := qs.Update(orm.Params{"value": value})
		if err != nil {
			o.Rollback()
			beego.Error(err)
			return err
		}
	} else {
		return errors.New("系统参数不存在")
	}
	o.Commit()
	return nil

}

//查询系统参数
func QueryParam(key string) (string, error) {
	o := orm.NewOrm()
	m := make([]*Param, 0)
	_, err := o.QueryTable(new(Param)).Filter("key", key).All(&m)
	if err != nil {
		return "", err
	}
	if len(m) == 0 {
		return "", errors.New("系统参数获取异常")
	}
	return m[0].Value, nil
}
