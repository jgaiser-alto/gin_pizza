package pizzas

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h handler) GetPizza(ctx *gin.Context) {
	var id, parseError = uuid.Parse(ctx.Param("id"))
	if parseError != nil {
		ctx.AbortWithError(http.StatusNotFound, parseError)
	}

	result, err := h.Repository.Get(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &result)
}
