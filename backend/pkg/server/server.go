package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "0.0.0",
		})
	})
}
