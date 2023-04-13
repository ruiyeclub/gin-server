package main

import (
	"gin-server/src/api"
	"gin-server/src/middleware"
	"gin-server/src/utils/config"
	"gin-server/src/utils/errors"
	"gin-server/src/utils/mylogs"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin日志处理
	f, err := mylogs.GetLogFile()
	if err == nil {
		gin.ForceConsoleColor()
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	r := gin.Default()

	// 统一异常处理
	r.Use(errors.Recover)

	// 跨域处理
	r.Use(middleware.Cors())

	// api routes
	api.UserAccountApi(r)

	port := config.Config.App.Port
	port = ":" + port
	r.Run(port)
}
