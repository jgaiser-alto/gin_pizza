package pizzas

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza/pkg/common/models"
)

func (h handler) GetPizza(ctx *gin.Context) {
	id := ctx.Param("id")

	var pizza models.Pizza

	if result := h.DB.First(&pizza, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &pizza)
}
