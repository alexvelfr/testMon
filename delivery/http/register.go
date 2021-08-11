package blockhttp

import (
	"github.com/alexvelfr/tmp/testmon"
	"github.com/gin-gonic/gin"
)

func RegisterBlockEndpoints(router *gin.RouterGroup, uc testmon.Usecase) {
	handler := newHandler(uc)
	router.POST("hoock_block", handler.UpdateBlock)
}
