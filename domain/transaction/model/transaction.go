package model

import (
	model2 "Test_MNC_2/domain/merchant/model"
	"Test_MNC_2/domain/user/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID         string
	Amount     int64
	UserId     string
	User       model.User
	MerchantId string
	Merchant   model2.Merchant
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.ID = uuid.New().String()
	return
}
