package repository

import (
	model2 "Test_MNC_2/domain/merchant/model"
	"gorm.io/gorm"
)

type MerchantRepositoryInterface interface {
	FirstByID(id string) (model2.Merchant, error)
	UpdateBalance(id string, amount int64) error
}

type merchantRepository struct {
	DB *gorm.DB
}

func MerchantRepository(DB *gorm.DB) MerchantRepositoryInterface {
	return &merchantRepository{
		DB: DB,
	}
}

func (mr *merchantRepository) UpdateBalance(id string, amount int64) error {
	merchant, _ := mr.FirstByID(id)
	merchant.Balance += amount
	return mr.DB.Save(&merchant).Error
}

func (mr *merchantRepository) FirstByID(id string) (model2.Merchant, error) {
	var merchant model2.Merchant
	if err := mr.DB.Where("id = ?", id).First(&merchant).Error; err != nil {
		return merchant, err
	}
	return merchant, nil
}
