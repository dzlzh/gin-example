package routes

import (
	"net/http"

	"github.com/dzlzh/gin-example/utils"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
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
