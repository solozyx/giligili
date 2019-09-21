package server

import (
	"giligili/api"
	"giligili/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	// 获取用户状态
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed := v1.Group("/")
		// 登录检查中间件
		authed.Use(middleware.AuthRequired())
		{
			// 通过中间件检查
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
			authed.POST("videos", api.CreateVideo)
			authed.PUT("video/:id", api.UpdateVideo)
			authed.DELETE("video/:id", api.DeleteVideo)
		}

		// 视频操作
		// v1.POST("videos", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		// v1.PUT("video/:id", api.UpdateVideo)
		// v1.DELETE("video/:id", api.DeleteVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 其他
		v1.POST("upload/token", api.UploadToken)
	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	r.StaticFile("/swagger.json", "./swagger/swagger.json")
	r.Static("/swagger", "./swagger/dist")

	return r
}
