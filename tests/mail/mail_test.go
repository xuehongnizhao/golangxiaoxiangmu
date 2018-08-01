package mail

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

func Test_Mail(t *testing.T) {
	initdatabase.InitMyDB(t)
	QueryMails(t)
}

//创建时间：2018年5月4日
//创建人：路锐
func QueryMails(t *testing.T) {
	opt := new(models.MailOption)
	qo := new(base.QueryOptions)
	qo.Limit = 10
	qo.Offset = 0
	opt.BaseOption = qo
	_, mail, err := QueryMail(opt)
	if err != nil {
		t.Fatal("查看用餐人员失败,失败信息为：", err)
	} else {
		if len(mail) == 0 {
			t.Log("查看用餐人员成功,但没有查询到符合要求的数据！")
		} else {
			t.Log("查看用餐人员成功")
			for _, v := range mail {
				t.Log(v)
			}
		}
	}
}
