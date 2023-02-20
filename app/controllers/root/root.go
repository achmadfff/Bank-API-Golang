package root

import (
	"Test_MNC_2/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	response.Json(context, http.StatusOK, nil, "OK")
}
