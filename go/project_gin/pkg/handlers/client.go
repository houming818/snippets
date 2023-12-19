package handlers

import (
	"github.com/gin-gonic/gin"
)

// GetClientIP 获取用户实际IP
func ClientIP(ctx *gin.Context) {
	Logger.Info("ClientIP %s", ctx.ClientIP())
	ctx.Writer.Write([]byte(ctx.ClientIP()))
	return
}
