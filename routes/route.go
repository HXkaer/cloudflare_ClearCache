package routes

import (
	"cloudflare_ClearCache/apiv1"
	"cloudflare_ClearCache/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式-没有输出
	}
	// 路由

	//r := gin.New() //终端不显示
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 首页
	r.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})

	// 注册路由
	v1 := r.Group("/api/v1")

	// post获取清理缓存的域名
	v1.POST("/domain", apiv1.DomainPost)

	// 无路由返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
