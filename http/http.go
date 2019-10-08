package http

import (
    "github.com/gin-gonic/gin"
    "io"
    "learnku-api/config"
    "learnku-api/middleware"
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
        root := api.Group("/")
        {
            root.POST("/register", userStore)
            root.GET("/refresh_token", userRefreshToken)
            root.POST("/login", userLogin)
        }

        // users
        users := api.Group("/users")
        // don't need auth middleware
        {
            //
        }
        // auth middleware
        users.Use(middleware.JWT())
        {
            users.GET("/", userList)
        }

        // topics
        topics := api.Group("/topics")
        // don't need auth middleware
        {
            topics.GET("/", topicsList)
            topics.GET("/:id", topicsById)
        }
        // auth middleware
        topics.Use(middleware.JWT())
        {
            // topics.POST("/store", topicsStore)
        }

        // categories
        categories := api.Group("/categories")
        {
            categories.GET("/", categoriesList)
        }
        categories.Use(middleware.JWT())
        {
            categories.POST("/store", categoriesStore)
        }
    }
}
