package pizzas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetPizzas(ctx *gin.Context) {

	result, err := h.Repository.GetAll()
	if err != nil {
		fmt.Printf("failed to get pizzas: %s\n", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &result)
}
