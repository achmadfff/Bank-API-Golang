package collection

import (
	"Test_MNC_2/config/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Test_MNC_2/app/controllers/user"
)

func MainRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRoleCtrl := user.RoleController(db)
	userRoute := main.Group("role")
	{
		userRoute.GET("/", userRoleCtrl.GetRole)
	}

	userAuthCtrl := user.AuthController(db)
	authRoute := main.Group("auth")
	{
		authRoute.POST("/register", userAuthCtrl.Register)
		authRoute.POST("/login", userAuthCtrl.Login)
		authRoute.POST("/logout", middleware.Auth(db), userAuthCtrl.Logout)
	}
	userTransCtrl := user.TransactionController(db)
	transRoute := main.Group("payment")
	{
		transRoute.POST("/transaction", middleware.Auth(db), userTransCtrl.CreateTransaction)
		transRoute.GET("/report", userTransCtrl.ReportTransaction)
	}
}
