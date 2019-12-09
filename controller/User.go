package controller

import (
	"gin-demo/core"
	"gin-demo/middleware"
	"gin-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// login
func Login(c *gin.Context) {
	middleware.JwtInstance().LoginHandler(c)
}

// refresh token
func RefreshToken(c *gin.Context) {
	middleware.JwtInstance().RefreshHandler(c)
}

// User list
func UserList(c *gin.Context) {
	q := core.DbInstance().GetDB().NewQuery("SELECT uid,nickname FROM `u_info` LIMIT 10")
	var user []model.User
	err := q.All(&user)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
