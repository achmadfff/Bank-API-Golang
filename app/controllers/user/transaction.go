package user

import (
	"Test_MNC_2/domain/merchant"
	repoMerchant "Test_MNC_2/domain/merchant/repository"
	"Test_MNC_2/domain/transaction"
	"Test_MNC_2/domain/transaction/model"
	"Test_MNC_2/domain/transaction/repository"
	user "Test_MNC_2/domain/user"
	repoUser "Test_MNC_2/domain/user/repository"
	"Test_MNC_2/lib/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type transactionController struct {
	TransactionService transaction.TransactionServiceInterface
}

func TransactionController(db *gorm.DB) *transactionController {
	return &transactionController{
		TransactionService: transaction.TransactionService(repository.TransactionRepository(db), merchant.MerchantService(repoMerchant.MerchantRepository(db)), user.AuthService(repoUser.AuthRepository(db))),
	}
}

func (tc *transactionController) CreateTransaction(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	var reqBody model.ReqBodyTransaction
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := tc.TransactionService.CreateTransaction(&reqBody, token); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, nil, "Success Create Transaction")
}

func (tc *transactionController) ReportTransaction(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	data, totalRow, errStatus, err := tc.TransactionService.GetReport(offset, limit)
	if err != nil {
		response.Error(ctx, errStatus, err.Error())
		return
	}

	response.JsonPagination(ctx, http.StatusOK, data, page, limit, totalRow)
}
