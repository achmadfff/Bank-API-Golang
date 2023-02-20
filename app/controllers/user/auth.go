package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"Test_MNC_2/domain/user"
	"Test_MNC_2/domain/user/model"
	"Test_MNC_2/domain/user/repository"
	"Test_MNC_2/lib/response"
)

type authController struct {
	AuthService user.AuthServiceInterface
}

func AuthController(db *gorm.DB) *authController {
	return &authController{
		AuthService: user.AuthService(repository.AuthRepository(db)),
	}
}

func (ac *authController) Register(context *gin.Context) {
	var reqBody model.ReqBodyRegister
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	if errStatus, err := ac.AuthService.Register(&reqBody); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, nil, "Success Register Customer")
}

func (ac *authController) Login(context *gin.Context) {
	var reqBody model.ReqBodyLogin
	if err := context.ShouldBind(&reqBody); err != nil {
		response.Error(context, http.StatusBadRequest, err.Error())
		return
	}
	resBody, errStatus, err := ac.AuthService.Login(&reqBody)
	if err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, resBody, "Success Login")
}

func (ac *authController) Logout(context *gin.Context) {
	user, err := ac.AuthService.CheckAuth(context.Request.Header.Get("Authorization"))
	if err != nil {
		response.Error(context, http.StatusUnauthorized, err.Error())
		return
	}
	if errStatus, err := ac.AuthService.Logout(*user); err != nil {
		response.Error(context, errStatus, err.Error())
		return
	}
	response.Json(context, http.StatusOK, nil, "Success Logout")
}
