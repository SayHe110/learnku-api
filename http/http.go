package http

import (
    "github.com/gin-gonic/gin"
    "io"
    "learnku-api/config"
    "learnku-api/pkg/response"
    "log"
    "os"
)

const (
    APIRoot     = "/api"
    LogFilePath = "./runtime/log/"
)

func Init() {
    file, _ := os.Create(LogFilePath + "gin.log")
    gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

    engine := gin.Default()

    gin.SetMode(config.AppConfig.Runmode)

    initRouter(engine)

    if err := engine.Run(config.AppConfig.Url); err != nil {
        log.Printf("engine.Start error(%v)", err)
        panic(err)
    }
}

func initRouter(e *gin.Engine) {
    api := e.Group(APIRoot)
    {
        users := api.Group("/users")
        {
            users.GET("/", func(context *gin.Context) {
                response.JSON(context, 200, "hello", nil)
            })
        }
    }
}
