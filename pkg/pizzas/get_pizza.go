package pizzas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pizza/pkg/common/models"
)

func (h handler) GetPizza(ctx *gin.Context) {
	var id, err = uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	var result *models.Pizza
	result, err = h.Repository.Get(id)
	if err != nil {
		fmt.Printf("failed to get pizza: %s\n", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &result)
}
