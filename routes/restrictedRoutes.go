package routes

import (
	"affordmed-cartservice/handlers"

	"github.com/gin-gonic/gin"
)

func RestrictedRoutes(res *gin.RouterGroup) {
	res.POST("/add-to-cart", handlers.AddToCart())
	res.POST("/remove", handlers.RemoveFromCart())
}
