package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

func topicsList(ctx *gin.Context) {
    res, _ := topicsSvc.GetTopicsList()
    response.JSON(ctx, http.StatusOK, "获取成功", res)
}

func topicsById(ctx *gin.Context) {
    res, _ := topicsSvc.GetTopicsById(ctx.Param("id"))
    response.JSON(ctx, http.StatusOK, "获取成功", res)
}
