package pizzas

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza/pkg/common/models"
)

type addPizzaRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h handler) AddPizza(ctx *gin.Context) {
	body := addPizzaRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pizza := models.Pizza{
		Name:        body.Name,
		Description: body.Description,
	}
	result, err := h.Repository.Create(pizza)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
