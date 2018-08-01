package order

import (
	"common/base"
	"supermarket/models"
	"supermarket/tests/initdatabase"
	"testing"
)

var id int64

func Test_Order(t *testing.T) {
	initdatabase.InitMyDB(t)
	AddOrder(t)
	UpdateOrders(t)
	QueryOrders(t)
	GetDetail(t)

}

//添加订单
func AddOrder(t *testing.T) {
	Details := make([]*models.Details, 0)
	details := new(models.Details)
	details.Ctime = base.GetCurrentData()
	details.GoodsName = base.GetUUID()
	details.GoodsNum = 50
	details.Price = 2.5
	Details = append(Details, details)
	order := new(models.Order)
	order.RtxAccount = base.GetUUID()
	order.Totalprice = float64(details.GoodsNum) * details.Price
	order.Ctime = base.GetCurrentData()
	order.Ptime = base.GetCurrentData()
	order.Status = base.GetUUID()
	err := InsertOrder(order, Details)
	if err != nil {
		t.Fatal("添加订单失败,失败信息为：", err)
	} else {
		id = Details[0].Id
		t.Log("添加订单成功")
	}
}

//查询订单
func QueryOrders(t *testing.T) {
	opt := new(models.OrderOption)
	qo := new(base.QueryOptions)
	qo.Limit = 10
	qo.Offset = 0
	opt.BaseOption = qo
	_, order, err := QueryOrder(opt)
	if err != nil {
		t.Fatal("查询订单失败,失败信息为：", err)
	} else {
		if len(order) == 0 {
			t.Log("查询订单成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询订单成功,查到的信息为：")
			for _, v := range order {
				t.Log(v)
			}
		}
	}
	opt.Id = -1
	_, order, err = QueryOrder(opt)
	if err != nil {
		t.Fatal("查询订单失败,失败信息为：", err)
	} else {
		if len(order) == 0 {
			t.Log("查询订单成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询订单成功,查到的信息为：")
			for _, v := range order {
				t.Log(v)
			}
		}
	}
}

//查询订单详细信息
func GetDetail(t *testing.T) {
	detail, err := GetDetails(id)
	if err != nil {
		t.Fatal("查询订单详细信息失败,失败信息为：", err)
	} else {
		if len(detail) == 0 {
			t.Log("查询订单详细信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询订单详细信息成功,查到的信息为：")
			for _, v := range detail {
				t.Log(v)
			}
		}
	}
	detail, err = GetDetails(-1)
	if err != nil {
		t.Fatal("查询订单详细信息失败,失败信息为：", err)
	} else {
		if len(detail) == 0 {
			t.Log("查询订单详细信息成功，但没有查询到符合要求的数据")
		} else {
			t.Log("查询订单详细信息成功,查到的信息为：")
			for _, v := range detail {
				t.Log(v)
			}
		}
	}
}

//修改订单信息状态
func UpdateOrders(t *testing.T) {
	order := new(models.Order)
	order.Status = base.GetUUID()
	err := UpdateOrderStatus(order)
	if err != nil {
		t.Fatal("修改订单信息状态失败,失败信息为：", err)
	} else {
		t.Log("修改订单信息状态成功")
	}
}
