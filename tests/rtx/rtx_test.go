package rtx

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

var id int64

func Test_Rtx(t *testing.T) {
	initdatabase.InitMyDB(t)
	GetRtxAccounts(t)
	ModifyTokens(t)
}

//创建时间：2018年5月4日
//创建人：路锐
func GetRtxAccounts(t *testing.T) {
	token := "0c13425328134dabbab4ce90942e0049"
	rtx, err := GetRtxAccount(token)
	if err != nil {
		t.Fatal("查看token失败,失败信息为：", err)
	} else {
		if len(rtx) == 0 {
			t.Log("查看token成功,但没有查询到符合要求的数据！")
		} else {
			t.Log("查看token成功")
			for _, v := range rtx {
				t.Log(v)
				id = v.Id
			}
		}
	}
}

//创建时间：2018年5月4日
//创建人：路锐
func ModifyTokens(t *testing.T) {
	ra := new(models.RtxAccount)
	ra.Id = id
	ra.Token = base.GetCurrentData()
	ra.Utime = base.GetCurrentData()
	err := ModifyToken(ra)
	if err != nil {
		t.Fatal("修改rtx失败,失败信息为：", err)
	} else {
		t.Log("修改rtx成功")
	}
}
