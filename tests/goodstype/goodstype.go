package goodstype

import (
	"supermarket/models"
)

//添加商品类型
func InsertGoodstype(gt *models.Goodstype) (*models.Goodstype, error) {
	gt, err := models.InsertGoodstype(gt)
	return gt, err
}

//查询商品类型
func QueryGoodstype(opt *models.GoodstypeOption) (int, []*models.Goodstype, error) {
	num, gt, err := models.QueryGoodstype(opt)
	return num, gt, err
}

//修改商品类型
func UpdateGoodstype(gt *models.Goodstype) (*models.Goodstype, error) {
	gt, err := models.UpdateGoodstype(gt)
	return gt, err
}

//删除
func DelGoodstype(id int64) (*models.Goodstype, error) {
	gt, err := models.DelGoodstype(id)
	return gt, err
}
