package http

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "io"
    "learnku-api/config"
    "learnku-api/middleware"
    categoryService "learnku-api/service/categories"
    topicService "learnku-api/service/topics"
    userService "learnku-api/service/users"
    "log"
    "os"
)

const (
    APIRoot     = "/api"
    LogFilePath = "./runtime/log/"
)

var (
    userSvc   *userService.Service
    topicsSvc *topicService.Service
    categorySvc *categoryService.Service
)

func Init(c *config.Config) {
    file, _ := os.Create(LogFilePath + "gin.log")
    gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

    engine := gin.Default()

    // session
    sessionStore := cookie.NewStore([]byte("secret"))
    engine.Use(sessions.Sessions(config.C.AppConfig.Name, sessionStore))

    gin.SetMode(config.C.AppConfig.Runmode)

    initService(c)
    initRouter(engine)

    if err := engine.Run(config.C.AppConfig.Url); err != nil {
        log.Printf("engine.Start error(%v)", err)
        panic(err)
    }
}

func initService(c *config.Config) {
    userSvc = userService.New(c)
    topicsSvc = topicService.New(c)
    categorySvc = categoryService.New(c)
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
        // users.Use(middleware.JWT())
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
        categories.Use(middleware.AuthSessionMiddle())
        {
            categories.POST("/store", categoriesStore)
        }
    }
}
