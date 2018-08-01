package order

import (
	"supermarket/models"
)

//查询订单
func QueryOrder(opt *models.OrderOption) (int, []*models.Order, error) {
	num, order, err := models.QueryOrder(opt)
	return num, order, err
}

//添加订单
func InsertOrder(order *models.Order, Details []*models.Details) error {
	err := models.InsertOrder(order, Details)
	return err
}

//查询订单详细信息
func GetDetails(order_id int64) ([]*models.Details, error) {
	details, err := models.GetDetails(order_id)
	return details, err
}

//修改订单信息状态
func UpdateOrderStatus(order *models.Order) error {
	err := models.UpdateOrderStatus(order)
	return err
}
