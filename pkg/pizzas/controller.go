package pizzas

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	Repository Repository
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := handler{
		Repository: CreateRepository(db),
	}
	routes := router.Group("/pizzas")
	routes.GET("", h.GetPizzas)
	routes.GET("/:id", h.GetPizza)
	routes.POST("", h.AddPizza)
	routes.PUT("/:id", h.UpdatePizza)
	routes.DELETE("/:id", h.DeletePizza)
}
