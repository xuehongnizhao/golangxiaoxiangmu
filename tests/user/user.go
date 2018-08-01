package user

import (
	"supermarket/models"
)

//查询用户信息
func QueryUserInfo(opt *models.QueryUserOptions) (int, []*models.User, error) {
	num, user, err := models.QueryUserInfo(opt)
	return num, user, err
}

func GetUser(account, password string) (*models.SUser, error) {
	su, err := models.GetUser(account, password)
	return su, err
}
