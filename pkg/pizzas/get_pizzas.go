package pizzas

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza/pkg/common/models"
)

func (h handler) GetPizzas(ctx *gin.Context) {
	var pizzas []models.Pizza

	if result := h.DB.Find(&pizzas); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &pizzas)
}
