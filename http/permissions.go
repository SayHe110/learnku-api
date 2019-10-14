package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

func permissionList(ctx *gin.Context) {
    //res,err := casbin.CasbinEnforce.GetRolesForUser("1231312a@a.c")
    //if err != nil {
    //    response.JSON(ctx, http.StatusInternalServerError, "获取失败", err)
    //}
    //log.Println(res)
    response.JSON(ctx, http.StatusOK, "成功", gin.H{
        "g": 123,
    })
}
