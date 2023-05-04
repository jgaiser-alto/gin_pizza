package pizzas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza/pkg/common/models"
)

type AddPizzaRequestBody struct {
	Name        string `json:"name"  binding:"required"`
	Description string `json:"description"  binding:"required"`
}

func (h handler) AddPizza(ctx *gin.Context) {
	body := AddPizzaRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.Repository.Create(toPizza(body))
	if err != nil {
		fmt.Printf("failed to create pizza: %s\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("suscessfully created pizza: %s\n", result)
	ctx.JSON(http.StatusCreated, result)
}
func toPizza(body AddPizzaRequestBody) models.Pizza {
	return models.Pizza{
		Name:        body.Name,
		Description: body.Description,
	}
}
