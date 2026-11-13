package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Version 由构建时 -ldflags 注入，默认 dev。
var Version = "dev"

// Health 返回服务健康状态，供探活使用。
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "web-service",
		"version": Version,
	})
}
