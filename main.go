package main

import (
	"os"

	"github.com/dzlzh/gin-example/models"
	"github.com/dzlzh/gin-example/routes"
	"github.com/dzlzh/gin-example/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	utils.RootPath, _ = os.Getwd()
	utils.InitConfig()

	// 初始化Redis
	utils.InitRedis()

	// 初始化DB
	models.InitDB()

	// 设置环境
	gin.SetMode(utils.Config.Section("system").Key("gin_mode").String())

	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// gin.DefaultErrorWriter = os.Stderr
	r := routes.SetupRouter()

	port := utils.Config.Section("system").Key("http_port").String()
	r.Run(":" + port)
}
