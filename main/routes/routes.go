package routes

import (
	"github.com/gin-gonic/gin"
	"reservas/main/controller"
)

func Routes(routes *gin.Engine) {
	tableGroup := routes.Group("/tables")
	tableGroup.GET("/", controller.FindAllTables)
	tableGroup.GET("/:id", controller.FindOneTables)
	tableGroup.PATCH("/:id", controller.PatchTables)
	tableGroup.DELETE("/:id", controller.DeleteTables)
	tableGroup.POST("/create", controller.CreateTables)
}
