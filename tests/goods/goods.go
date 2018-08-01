package goods

import (
	"supermarket/models"
)

//查询商品信息
func QueryGood(opt *models.QueryGoods) ([]*models.Goods, int, error) {
	goods, num, err := models.QueryGood(opt)
	return goods, num, err
}

//添加商品信息
func AddGoods(goods *models.Goods) (*models.Goods, error) {
	goods, err := models.AddGoods(goods)
	return goods, err
}

//修改商品信息
func UpdateGoods(goods *models.Goods) (*models.Goods, error) {
	goods, err := models.UpdateGoods(goods)
	return goods, err
}

//删除商品信息
func DeleteGoods(id int64) (*models.Goods, error) {
	goods, err := models.DeleteGoods(id)
	return goods, err
}
