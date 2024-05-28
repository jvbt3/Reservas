package server

import (
	"github.com/gin-gonic/gin"
	"reservas/main/routes"
)

func Server() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r = SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	routes.Routes(router)

	return router
}
