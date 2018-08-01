package goods

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

var id int64

func Test_Goods(t *testing.T) {
	initdatabase.InitMyDB(t)
	InsertGoods(t)
	UpdaGoods(t)
	DeleGoods(t)
	QueryGoods(t)
}

//添加商品信息
func InsertGoods(t *testing.T) {
	goods := new(models.Goods)
	goods.Name = base.GetUUID()
	goods.Price = 1.5
	goods.Total = 50

	goods.Ctime = base.GetCurrentData()
	goods.Utime = base.GetCurrentData()
	goods.Goodstype = base.GetUUID()
	goods, err := AddGoods(goods)
	if err != nil {
		t.Fatal("添加商品失败,失败信息为：", err)
	} else {
		id = goods.Id
		t.Log("添加商品成功,新增的信息为：", goods)
	}
}

//修改商品信息
func UpdaGoods(t *testing.T) {
	goods := new(models.Goods)
	goods.Id = id
	goods.Name = base.GetUUID()
	goods.Price = 2.5
	goods.Utime = base.GetCurrentData()
	goods, err := UpdateGoods(goods)
	if err != nil {
		t.Fatal("修改商品失败,失败信息为：", err)
	} else {
		if goods == nil {
			t.Log("修改商品成功,但要修改的数据不存在")
		} else {
			t.Log("修改商品成功,修改的信息为：", goods)
		}
	}
}

//删除商品信息
func DeleGoods(t *testing.T) {
	goods, err := DeleteGoods(138)
	if err != nil {
		t.Fatal("删除商品失败,失败信息为：", err)
	} else {
		if goods == nil {
			t.Log("删除商品成功,但要删除的数据不存在")
		} else {
			t.Log("删除商品成功,删除的信息为：", goods)
		}
	}
}

//查询商品信息
func QueryGoods(t *testing.T) {
	opt := new(models.QueryGoods)
	qo := new(base.QueryOptions)
	qo.Limit = 10
	qo.Offset = 0
	opt.BaseOption = qo
	opt.Status = "enabled"
	goods, _, err := QueryGood(opt)
	if err != nil {
		t.Fatal("查询商品信息失败,失败信息为：", err)
	} else {
		if len(goods) == 0 {
			t.Log("查询商品信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询商品信息成功,查到的信息为：")
			for _, v := range goods {
				t.Log(v)
			}
		}
	}
	opt.Name = "?"
	goods, _, err = QueryGood(opt)
	if err != nil {
		t.Fatal("查询商品信息失败,失败信息为：", err)
	} else {
		if len(goods) == 0 {
			t.Log("查询商品信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询商品信息成功,查到的信息为：")
			for _, v := range goods {
				t.Log(v)
			}
		}
	}
	opt.Name = "奶"
	goods, _, err = QueryGood(opt)
	if err != nil {
		t.Fatal("查询商品信息失败,失败信息为：", err)
	} else {
		if len(goods) == 0 {
			t.Log("查询商品信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询商品信息成功,查到的信息为：")
			for _, v := range goods {
				t.Log(v)
			}
		}
	}
}
