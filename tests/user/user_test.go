package user

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

var id int64

func Test_User(t *testing.T) {
	initdatabase.InitMyDB(t)
	QueryUserInfos(t)
	GetUsers(t)

}

//查询用户信息
//创建时间：2018年5月4日
//创建人：路锐
func QueryUserInfos(t *testing.T) {
	opt := new(models.QueryUserOptions)
	qo := new(base.QueryOptions)
	qo.Limit = 10
	qo.Offset = 0
	opt.BaseOptions = qo
	_, user, err := QueryUserInfo(opt)
	if err != nil {
		t.Fatal("查询用户信息失败,失败信息为：", err)
	} else {
		if len(user) == 0 {
			t.Log("查询用户信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询用户信息成功,查到的信息为：")
			for _, v := range user {
				t.Log(v)
			}
		}
	}
	opt.Name = "?"
	_, user, err = QueryUserInfo(opt)
	if err != nil {
		t.Fatal("查询用户信息失败,失败信息为：", err)
	} else {
		if len(user) == 0 {
			t.Log("查询用户信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询用户信息成功,查到的信息为：")
			for _, v := range user {
				t.Log(v)
			}
		}
	}
}

func GetUsers(t *testing.T) {
	account := "123456"
	password := "123456"
	su, err := GetUser(account, password)
	if err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			t.Log("查询用户信息成功,但是查询不到符合要求的数据！")
		} else {
			t.Fatal("查询用户信息失败", err)
		}

	} else {
		if su == nil {
			t.Log("查询用户信息成功,但是查询不到符合要求的数据！")
		} else {
			t.Log("查询用户信息成功,查到的信息为：", su)
		}
	}
}
