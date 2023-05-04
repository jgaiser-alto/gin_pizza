package pizzas

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *handler) DeletePizza(ctx *gin.Context) {
	var id, err = uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	if pizza, err := h.Repository.Get(id); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	} else {
		if err := h.Repository.Delete(*pizza); err != nil {
			fmt.Printf("failed to deleted pizza: %s error: %s\n", id, err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	fmt.Printf("suscessfully deleted pizza: %s", id)
	// Responds with no response body
	ctx.Status(http.StatusOK)
}
