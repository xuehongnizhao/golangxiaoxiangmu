package goodstype

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

var id int64

func Test_Goodstype(t *testing.T) {
	initdatabase.InitMyDB(t)
	AddGoodstype(t)
	UpdateGoodstypes(t)
	DeleteGoodsType(t)
	GetMenuType(t)
}

//添加商品类型
//创建时间：2018年5月4日
//创建人：路锐
func AddGoodstype(t *testing.T) {
	gt := new(models.Goodstype)
	gt.Typename = base.GetUUID()
	gt.Ctime = base.GetCurrentData()
	gt.Status = "enable"
	gt, err := InsertGoodstype(gt)
	if err != nil {
		t.Fatal("添加商品类型失败,失败信息为：", err)
	} else {
		id = gt.Id
		t.Log("添加商品类型成功,新增的信息为：", gt)
	}
}

//修改商品类型
//创建时间:2018年5月4日
//创建人：路锐
func UpdateGoodstypes(t *testing.T) {
	gt := new(models.Goodstype)
	gt.Id = id
	gt.Typename = base.GetUUID()
	gt, err := UpdateGoodstype(gt)
	if err != nil {
		t.Log("修改商品类型失败,失败信息为：", err)
	} else {
		if gt == nil {
			t.Log("修改商品类型成功，但修改的数据不存在")
		} else {
			t.Log("修改商品类型成功", gt)
		}
	}
}

//删除商品类型
//创建时间:2018年5月4日
//创建人：路锐
func DeleteGoodsType(t *testing.T) {
	gt, err := DelGoodstype(id)
	if err != nil {
		t.Fatal("删除商品类型失败,失败信息为：", err)
	} else {
		if gt == nil {
			t.Log("删除商品类型成功，要删除的信息不存在")
		} else {
			t.Log("删除商品类型成功,删除的信息为：", gt)
		}
	}
}

//查看商品类型
//创建时间:2018年5月3日
//创建人：路锐
func GetMenuType(t *testing.T) {
	opt := new(models.GoodstypeOption)
	qo := new(base.QueryOptions)
	qo.Limit = 10
	qo.Offset = 0
	opt.BaseOption = qo
	opt.Status = "enabled"
	_, gt, err := QueryGoodstype(opt)
	if err != nil {
		t.Fatal("查看商品类型失败,失败信息为：", err)
	} else {
		if len(gt) == 0 {
			t.Log("查看商品类型成功,但没有查询到符合要求的数据！")
		} else {
			t.Log("查看商品类型成功")
			for _, v := range gt {
				t.Log(v)
			}
		}
	}
	opt.Typename = "?"
	_, gt, err = QueryGoodstype(opt)
	if err != nil {
		t.Fatal("查看商品类型失败,失败信息为：", err)
	} else {
		if len(gt) == 0 {
			t.Log("查看商品类型成功,但没有查询到符合要求的数据！")
		} else {
			t.Log("查看商品类型成功")
			for _, v := range gt {
				t.Log(v)
			}
		}
	}
}
