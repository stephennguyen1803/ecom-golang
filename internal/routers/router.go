package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "pong ... ping" + name,
		"uid":     uid,
		"users":   []string{"cr7", "messi", "neymar"},
	})
}
