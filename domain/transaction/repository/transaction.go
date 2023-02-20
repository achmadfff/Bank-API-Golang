package repository

import (
	model2 "Test_MNC_2/domain/transaction/model"
	"errors"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	Create(reqBody *model2.ReqBodyTransaction, userId string) error
	GetAll(offset, limit int) ([]model2.ResBodyReportTransaction, int64, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func TransactionRepository(DB *gorm.DB) TransactionRepositoryInterface {
	return &transactionRepository{
		DB: DB,
	}
}

func (r *transactionRepository) Create(reqBody *model2.ReqBodyTransaction, userId string) error {
	var transaction model2.Transaction
	transaction.Amount = reqBody.Amount
	transaction.UserId = userId
	transaction.MerchantId = reqBody.MerchantId
	return r.DB.Create(&transaction).Error
}

func (r *transactionRepository) GetAll(offset, limit int) ([]model2.ResBodyReportTransaction, int64, error) {
	var transaction []model2.ResBodyReportTransaction
	var totalRow int64

	if err := r.DB.Table("transactions").Count(&totalRow).Offset(offset).Limit(limit).Scan(&transaction).Error; err != nil {
		return transaction, totalRow, errors.New("failed to get transaction : " + err.Error())
	}

	return transaction, totalRow, nil
}
