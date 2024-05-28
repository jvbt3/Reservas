package controller

import (
	"github.com/gin-gonic/gin"
	"reservas/main/service"
)

func CreateTables(c *gin.Context) {
	c.JSON(200, service.Create())
}

func FindAllTables(c *gin.Context) {
	c.JSON(200, service.FindAll())
}

func FindOneTables(c *gin.Context) {
	c.JSON(200, service.FindOne())
}

func PatchTables(c *gin.Context) {
	c.JSON(200, service.Patch())
}

func DeleteTables(c *gin.Context) {
	c.JSON(200, service.Delete())
}
