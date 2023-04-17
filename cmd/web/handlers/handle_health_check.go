package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pedidosya/@project_name@/engine"

)

func HandleHealthCheck(c *gin.Context, e engine.Spec) {
	c.JSON(200, gin.H{
		"env":     os.Getenv("ENV"),
		"name":    "@project_name@",
		"version": os.Getenv("VERSION"),
	})
}
