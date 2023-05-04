package pizzas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pizza/pkg/common/models"
)

type UpdatePizzaRequestBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h handler) UpdatePizza(ctx *gin.Context) {
	var id, err = uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	body := UpdatePizzaRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var pizza *models.Pizza
	pizza, err = h.Repository.Get(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	pizza.Name = body.Name
	pizza.Description = body.Description
	pizza, err = h.Repository.Update(*pizza)
	if err != nil {
		fmt.Printf("failed to update pizza: %s\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}
