package rtx

import (
	"supermarket/models"
)

func GetRtxAccount(token string) ([]*models.RtxAccount, error) {
	rtx, err := models.GetRtxAccount(token)
	return rtx, err
}

func ModifyToken(ra *models.RtxAccount) error {
	err := models.ModifyToken(ra)
	return err
}
