package routers

import (
	"fmt"
	"ginDemo/routers/api/v1"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"ginDemo/pkg/setting"
	"ginDemo/models"

)

func InitRouter() *gin.Engine {
	r := gin.New()
	var identityKey = "id"

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(setting.ENV),
		Timeout:     7 * 24 * time.Hour,
		MaxRefresh:  7 * 24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID, //传入的值 将在Authorizator中使用到
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := &models.User{}
			user.ID = int(claims["id"].(float64))
			return user
		},
		Authenticator: Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: AuthErrorHandler,

		TokenLookup: "header: token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		fmt.Println("JWT Error:" + err.Error())
	}



	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON( http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": "Page not found"})
	})

	r0 := r.Group("/login")
	{
		r0.POST("/getToken", authMiddleware.LoginHandler)//获取token
	}

	r1 := r.Group("/test")
	r1.Use(authMiddleware.MiddlewareFunc())
	{
		r1.POST("/index", v1.TestIndexHandler)

	}

	return r
}
//登录
func Login (c *gin.Context) (interface{}, error)  {
	user := models.User{
		UserName:"sss",
	}
	user.ID=1
	return user, nil
}

func AuthSuccessHandler(c *gin.Context, code int, token string, expire time.Time) {
	var tokenData = map[string]string{"token":token, "expire":expire.Format(time.RFC3339)}
	UserBase, _ := c.Get("UserBase")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "success",
		"data": tokenData,
		"userBase": UserBase,
	})
}

func AuthErrorHandler(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg": message,
	})
}
