package middleware

import (
	"affordmed-cartservice/routes"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare(router *gin.Engine) {

	restricted := router.Group("/r")
	routes.RestrictedRoutes(restricted)
}
