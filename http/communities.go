package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

func communityList(ctx *gin.Context) {
    res, err := communitiesSvc.GetList();
    if err != nil {
        response.JSON(ctx, http.StatusOK, "获取列表失败", err)
        return
    }

    response.JSON(ctx, http.StatusOK, "获取列表成功", res)
}

func communityUpdate(ctx *gin.Context) {

}
