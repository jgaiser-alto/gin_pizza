package pizzas

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UpdatePizzaRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *handler) UpdatePizza(ctx *gin.Context) {
	var id, parseError = uuid.Parse(ctx.Param("id"))
	if parseError != nil {
		ctx.AbortWithError(http.StatusNotFound, parseError)
	}
	body := UpdatePizzaRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pizza, err := h.Repository.Get(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	pizza.Name = body.Name
	pizza.Description = body.Description

	h.Repository.Update(*pizza)

	ctx.JSON(http.StatusOK, &pizza)
}
