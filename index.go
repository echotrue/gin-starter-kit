package main

import (
	"gin-demo/core"
	"gin-demo/core/config"
	"gin-demo/core/redis"
	"gin-demo/middleware"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// parse config file
	config.NewToml().Parse()
	c := config.GetConfig()

	gin.SetMode(c.GinModel)

	r := gin.New()

	// logger init
	core.New(r).FileTarget()

	// jwt init
	middleware.NewJWT()

	// Connect db
	core.DbInstance().NewDB()

	// Connect Redis
	redis.Instance().Connect()

	// Load template
	r.LoadHTMLGlob("./views/**/*")

	// Error handle
	r.Use(func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": c.Errors,
			})
			return
		}
		c.Next()
	})

	// register router
	router.Register(r)

	// validator translate
	//core.TsInstance().NewTs()

	_ = r.Run(":8888")
}
