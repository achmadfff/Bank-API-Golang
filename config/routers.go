package config

import (
	"Test_MNC_2/app/controllers/root"
	"Test_MNC_2/config/collection"
	"Test_MNC_2/db"
	"github.com/gin-gonic/gin"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)

	Routers.GET("/", root.Index)
	main := Routers.Group("v1")
	collection.MainRouter(db.DB, main)
}
