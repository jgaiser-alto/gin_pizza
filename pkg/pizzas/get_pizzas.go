package pizzas

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) GetPizzas(ctx *gin.Context) {

	result, err := h.Repository.GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &result)
}
