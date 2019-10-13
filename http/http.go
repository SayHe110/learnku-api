package http

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
    "io"
    "learnku-api/config"
    "learnku-api/middleware"
    "log"
    "os"

    categoryService "learnku-api/service/categories"
    communityService "learnku-api/service/communities"
    topicService "learnku-api/service/topics"
    userService "learnku-api/service/users"
)

const (
    APIRoot     = "/api"
    LogFilePath = "./runtime/log/"
)

var (
    userSvc        *userService.Service
    topicsSvc      *topicService.Service
    categorySvc    *categoryService.Service
    communitiesSvc *communityService.Service
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
    communitiesSvc = communityService.New(c)
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

        // communities
        communities := api.Group("/communities", middleware.AuthSessionMiddle("response", "write"))
        {
            communities.GET("", communityList)
            communities.PATCH("/update", communityUpdate)
        }

        // users
        users := api.Group("/users")
        {
            users.GET("/", userList)
        }
        // users.Use(middleware.AuthSessionMiddle())
        {
            users.GET("/logout", userLogout)
        }

        // topics
        topics := api.Group("/topics")
        {
            topics.GET("/", topicsList)
            topics.GET("/:id", topicsById)
        }
        topics.Use(middleware.JWT())
        {
            // topics.POST("/store", topicsStore)
        }

        // categories
        categories := api.Group("/categories")
        {
            categories.GET("", categoriesList)
        }
        // categories.Use(middleware.AuthSessionMiddle())
        {
            categories.POST("/store", middleware.AuthSessionMiddle("response", "write"), categoriesStore)
        }
    }
}
