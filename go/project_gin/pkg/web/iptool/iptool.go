package iptool

import (
	"github.com/gin-gonic/gin"
)

// GetClientIP 获取用户实际IP
func GetClientIP(ctx *gin.Context) {
	ctx.Writer.Write([]byte(ctx.ClientIP()))
	return
}
