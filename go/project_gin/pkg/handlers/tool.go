package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

// GetClientIP 获取用户实际IP
func Ping(ctx *gin.Context) {
	Logger.Info("Ping %s", ctx.ClientIP())
	hostname, err := os.Hostname()
	if err != nil {
		Logger.Error("Ping %s", err.Error())
		ctx.Writer.Write([]byte(err.Error()))
	} else {
		ctx.Writer.Write([]byte(hostname))
	}
	return
}
