package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

func permissionList(ctx *gin.Context) {

    response.JSON(ctx, http.StatusOK, "成功", nil)
}
