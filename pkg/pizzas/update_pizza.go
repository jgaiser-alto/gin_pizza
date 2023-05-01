package pizzas

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pizza/pkg/common/models"
)

type UpdatePizzaRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h handler) UpdatePizza(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdatePizzaRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pizza models.Pizza

	if result := h.DB.First(&pizza, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	pizza.Name = body.Name
	pizza.Description = body.Description

	h.DB.Save(&pizza)

	ctx.JSON(http.StatusOK, &pizza)
}
