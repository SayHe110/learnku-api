package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/config"
    "log"
)

const (
    APIRoot = "/api"
)

func Init() {
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
                context.JSON(200, gin.H{
                    "message": "Hello world",
                })
            })
        }
    }
}
