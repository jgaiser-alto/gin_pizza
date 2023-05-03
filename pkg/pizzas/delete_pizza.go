package pizzas

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *handler) DeletePizza(ctx *gin.Context) {
	var id, parseError = uuid.Parse(ctx.Param("id"))
	if parseError != nil {
		ctx.AbortWithError(http.StatusNotFound, parseError)
	}

	pizza, err := h.Repository.Get(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	h.Repository.Delete(*pizza)

	ctx.JSON(http.StatusOK, &pizza)
}
