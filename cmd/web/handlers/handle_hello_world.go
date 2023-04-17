package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pedidosya/@project_name@/engine"
)

func HandleHelloWorld(c *gin.Context, e engine.Spec) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": e.Hello(name),
	})
}
