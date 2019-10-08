package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    topicSvc "learnku-api/service/topics"
    "net/http"
)

func topicsList(ctx *gin.Context) {
    res, _ := topicSvc.Topics()
    response.JSON(ctx, http.StatusOK, "获取成功", res)
}

func topicsById(ctx *gin.Context) {
    res, _ := topicSvc.TopicsById(ctx.Param("id"))
    response.JSON(ctx, http.StatusOK, "获取成功", res)
}
