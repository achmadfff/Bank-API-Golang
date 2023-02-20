package transaction

import (
	"Test_MNC_2/domain/merchant"
	model2 "Test_MNC_2/domain/transaction/model"
	"Test_MNC_2/domain/transaction/repository"
	"Test_MNC_2/domain/user"
	"net/http"
)

type TransactionServiceInterface interface {
	CreateTransaction(reqBody *model2.ReqBodyTransaction, token string) (int, error)
	GetReport(offset, limit int) ([]model2.ResBodyReportTransaction, int64, int, error)
}

type transactionService struct {
	Repository      repository.TransactionRepositoryInterface
	MerchantService merchant.MerchantServiceInterface
	AuthService     user.AuthServiceInterface
}

func TransactionService(repository repository.TransactionRepositoryInterface, merchantService merchant.MerchantServiceInterface, authService user.AuthServiceInterface) TransactionServiceInterface {
	return &transactionService{
		Repository:      repository,
		MerchantService: merchantService,
		AuthService:     authService,
	}
}

func (ts *transactionService) CreateTransaction(reqBody *model2.ReqBodyTransaction, token string) (int, error) {
	user, errCheck := ts.AuthService.CheckAuth(token)
	if errCheck != nil {
		return http.StatusInternalServerError, errCheck
	}
	userId := user.ID
	err := ts.Repository.Create(reqBody, userId)
	errUpdate := ts.MerchantService.UpdateBalance(reqBody.MerchantId, reqBody.Amount)
	if errUpdate != nil {
		return http.StatusInternalServerError, errUpdate
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ts *transactionService) GetReport(offset, limit int) ([]model2.ResBodyReportTransaction, int64, int, error) {
	var roles []model2.ResBodyReportTransaction
	roles, totalRow, err := ts.Repository.GetAll(offset, limit)
	if err != nil {
		return roles, totalRow, http.StatusInternalServerError, err
	}

	return roles, totalRow, http.StatusOK, nil
}
