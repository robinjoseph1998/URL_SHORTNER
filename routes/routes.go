package routes

import (
	"URL_Shortner/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.POST("/shorter", handlers.URLshorter)
}
