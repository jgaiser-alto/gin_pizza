package pizzas

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *handler) DeletePizza(ctx *gin.Context) {
	var id, parseError = uuid.Parse(ctx.Param("id"))
	if parseError != nil {
		ctx.AbortWithError(http.StatusNotFound, parseError)
	}

	if pizza, getErr := h.Repository.Get(id); getErr != nil {
		ctx.AbortWithError(http.StatusNotFound, getErr)
	} else {
		if deleteErr := h.Repository.Delete(*pizza); deleteErr != nil {
			ctx.AbortWithError(http.StatusInternalServerError, deleteErr)
		}
	}

	// Responds with no response body
	ctx.Status(http.StatusOK)
}
