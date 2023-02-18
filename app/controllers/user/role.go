package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"Test_MNC_2/domain/user"
	"Test_MNC_2/domain/user/repository"
	"Test_MNC_2/lib/response"
)

type roleController struct {
	RoleService user.RoleServiceInterface
}

func RoleController(db *gorm.DB) *roleController {
	return &roleController{
		RoleService: user.RoleService(repository.RoleRepository(db)),
	}
}

func (rc *roleController) GetRole(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	data, totalRow, errStatus, err := rc.RoleService.GetRole(offset, limit)
	if err != nil {
		response.Error(ctx, errStatus, err.Error())
		return
	}

	response.JsonPagination(ctx, http.StatusOK, data, page, limit, totalRow)
}
