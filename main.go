package main

import (
	"URL_Shortner/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.File("assets/index.html")
	})
	routes.SetUpRoutes(router)
	router.Run(":7000")
}
