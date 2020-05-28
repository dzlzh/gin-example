package routes

import (
	"net/http"
	"time"

	"github.com/dzlzh/gin-example/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// 跨域
	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	Router.Use(mwCORS)
}

func SetupRouter() *gin.Engine {

	Router.StaticFile("favicon.ico", utils.RootPath+"/favicon.ico")

	Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "helloweb",
		})
	})

	Router.NoRoute(func(c *gin.Context) {
		code := http.StatusNotFound
		c.JSON(code, gin.H{
			"code":    code,
			"message": http.StatusText(code),
			"data":    gin.H{},
		})
	})

	return Router
}
