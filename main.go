package main

import (
	"URL_Shortner/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetUpRoutes(router)
	router.Run(":7000")

}
