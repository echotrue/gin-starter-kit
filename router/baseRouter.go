package router

import (
	"gin-demo/controller"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	//homepage
	r.Any("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index/index", nil)
	})

	// 404
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
		})
	})

	//v1
	v1 := r.Group("/v1")
	{
		// login
		v1.POST("/login", controller.Login)

		// redis
		v1.GET("/redis", controller.RdsIndex)
		v1.POST("/redis/select", controller.SelectDB)
		v1.POST("/redis/search", controller.Search)
		v1.POST("/redis/excuse", controller.Excuse)

		// Test
		v1.GET("/test", controller.IndexTest)
		// Authorize token
		auth := v1.Group("", middleware.JwtInstance().MiddlewareFunc())
		{
			// refresh token
			auth.POST("/refresh", controller.RefreshToken)
			// User list
			auth.GET("/user_list", controller.UserList)
		}
	}

}
