package router

import (
	"fmt"
	"gin-web-skeleton/app/admin/controller"
	"gin-web-skeleton/app/index/api"
	"gin-web-skeleton/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {

	g.Static("/static", "./public/static")
	// 格式化日志输出格式
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	g.Use(gin.Recovery())
	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The router is not exist!")
	})

	// 服务器健康检查
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g.GET("/", api.Hello)

	// 管理后台路由
	//LoadAdminRouter(g)

	r := g.Group("/admin", middleware.AdminAuth())
	{
		g.LoadHTMLGlob("app/admin/templates/**/*")

		r.GET("/login", controller.Login)
		r.POST("/login", controller.Login)
		r.GET("/index", controller.Index)
		r.POST("/index", controller.Index)
		r.GET("/welcome", controller.Welcome)
	}

	return g
}
