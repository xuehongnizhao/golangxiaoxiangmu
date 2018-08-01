package mail

import (
	"supermarket/models"
)

func QueryMail(opt *models.MailOption) (int, []*models.Mail, error) {
	num, mail, err := models.QueryMail(opt)
	return num, mail, err
}
