package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"lovelake.cn/app/pkg/web/iptool"
	"lovelake.cn/texiusi/pkg/config"
	"lovelake.cn/texiusi/pkg/web/render"
)

// Run 新建HTTP服务
func Run() error {
	if !config.GlobalConfig.HTTPConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.GlobalConfig.LogConfig.DisableColor {
		gin.DisableConsoleColor()
	}

	f, _ := os.Create(config.GlobalConfig.LogConfig.Filename)
	gin.DefaultWriter = io.MultiWriter(f)

	globalWeb := gin.New()
	pprof.Register(globalWeb)

	// 注册 Middlewares
	//	registerMiddlewares(web)

	// 注册 Hanlders
	registerHandlers(globalWeb)

	if err := globalWeb.Run(config.GlobalConfig.HTTPConfig.Host); err != nil {
		return err
	}
	return nil
}

func registerHandlers(web *gin.Engine) error {
	rg := web.Group(config.GlobalConfig.HTTPConfig.Prefix)

	rg.GET("/ip", iptool.GetClientIP)

	web.LoadHTMLGlob("web/templates/*")

	rg.GET("/html/:tmpl_name", func(ctx *gin.Context) {
		tmplName := ctx.Param("tmpl_name")
		r, err := render.RenderFactory(tmplName, ctx)
		if err != nil || r == nil {
			log.Println("render error: ", err)
			ctx.HTML(http.StatusInternalServerError, "error500.tmpl", nil)
			return
		}
		dataH, err := r.GetDataH(ctx)
		if err != nil {
			fmt.Printf("get data error: %v \n", err)
			ctx.HTML(http.StatusInternalServerError, tmplName+".tmpl", nil)
		}
		ctx.HTML(http.StatusOK, tmplName+".tmpl", dataH)
	})

	return nil

}
