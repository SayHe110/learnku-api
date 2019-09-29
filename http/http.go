package http

import (
    "github.com/gin-gonic/gin"
    "io"
    "learnku-api/config"
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

    // validator error handler
    //uni := ut.New(zh.New(), zh.New())
    //engine.Use(middleware.NewValidatorErrorHandler(uni).HandlerValidatorError)

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
            users.GET("/", userList)
            users.POST("/store", userStore)
        }
    }
}
