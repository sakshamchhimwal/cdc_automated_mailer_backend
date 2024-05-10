package utils

import (
	"cdc_mailer/models"
	"cdc_mailer/services"
	"errors"
	"gorm.io/gorm"
)

func VerifyUser(userId uint) error {
	var findUser models.User
	result := services.DB.First(&findUser, userId)

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func VerifyUserCompany(userId uint, companyId uint) (models.Company, error) {
	var findCompany models.Company
	result := services.DB.First(&findCompany).Where(&models.Company{HandlerID: userId, Model: gorm.Model{ID: companyId}})

	if result.RowsAffected == 0 {
		return findCompany, errors.New("company not handled by user")
	}

	return findCompany, nil
}
