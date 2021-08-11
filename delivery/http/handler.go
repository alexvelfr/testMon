package blockhttp

import (
	"github.com/alexvelfr/tmp/testmon"
	"github.com/gin-gonic/gin"
)

type response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}
type handler struct {
	uc testmon.Usecase
}

func (h *handler) UpdateBlock(ctx *gin.Context) {
	type inp struct {
		Name string `json:"name"`
	}
	i := inp{}
	ctx.BindJSON(&i)
	err := h.uc.UpdateBlock(i.Name)
	if err != nil {
		ctx.JSON(200, response{
			Status:      "error",
			Description: err.Error(),
		})
		return
	}
	ctx.JSON(200, response{
		Status:      "success",
		Description: "",
	})
}

func newHandler(uc testmon.Usecase) *handler {
	return &handler{
		uc: uc,
	}
}
