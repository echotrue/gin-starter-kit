package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var (
	identityKey = "id"
	err         error
	jwtInstance *jwt.GinJWTMiddleware
)

// Get jwt instance
func JwtInstance() *jwt.GinJWTMiddleware {
	return jwtInstance
}

// Init jwt config
func NewJWT() {
	// jwt middleware
	jwtInstance, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		//SigningAlgorithm:      "", //签名算法，默认 HS256
		Key:        []byte("secret key"),
		Timeout:    time.Hour, //token 有效期
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var loginVals login
			loginVals.Username = c.PostForm("username")
			loginVals.Password = c.PostForm("password")

			if errs := c.ShouldBind(&loginVals); errs != nil {

				// todo
				//for _, err := range errs.(validator.ValidationErrors) {

				/*fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())*/

				//}
				return nil, errs
				//return nil, jwt.ErrMissingLoginValues
			}

			userId := loginVals.Username
			password := loginVals.Password

			// check password
			if userId == "admin" && password == "admin" {
				return &User{
					UserName:  "admin",
					FirstName: "huang",
					LastName:  "axl",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(context *gin.Context, code int, message string) {
			context.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 登陆成功后响应数据格式
		LoginResponse: func(context *gin.Context, i int, s string, i2 time.Time) {
			context.JSON(http.StatusOK, gin.H{
				"code":         i,
				"access_token": s,
				"expired_at":   i2.Format("2006-01-02 15:04:05"),
			})
		},
		//刷新token响应数据格式
		RefreshResponse: func(context *gin.Context, i int, s string, i2 time.Time) {
			context.JSON(http.StatusOK, gin.H{
				"code":         i,
				"access_token": s,
				"expired_at":   i2.Format("2006-01-02 15:04:05"),
			})
		},
		IdentityHandler: func(context *gin.Context) interface{} {
			claims := jwt.ExtractClaims(context)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			return e.Error()
		},
		//PrivKeyFile:       "",
		//PubKeyFile:        "",
		SendCookie: true,
		//SecureCookie:      false,
		//CookieHTTPOnly:    false,
		//CookieDomain:      "",
		//SendAuthorization: false,
		//DisabledAbort:     false,
		CookieName: "GIN-API-TOKEN",
	})

	if err != nil {
		//panic("JWT Error:" + err.Error())
		log.Fatal("JWT Error:" + err.Error())
	}
}
