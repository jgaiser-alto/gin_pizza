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

	var pizza models.Pizza

	pizza.Name = body.Name
	pizza.Description = body.Description

	if result := h.DB.Create(&pizza); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &pizza)
}
