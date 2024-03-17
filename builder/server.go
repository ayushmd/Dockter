package builder

import (
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", pong)
	// other handlers
	return r
}
