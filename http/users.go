package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/model/users"
    "learnku-api/pkg/response"
    userSvc "learnku-api/service/users"
    "net/http"
)

func userList(c *gin.Context) {
    res, _ := userSvc.Users()
    response.JSON(c, http.StatusOK, "获取成功", res)
}

func userStore(c *gin.Context) {

    userRes := &users.UserStoreParam{}

    if err := c.ShouldBind(&userRes); err != nil {
        response.JSON(c, 50001, "参数绑定错误", nil)
        return
    }

    if err := userSvc.Store(userRes); err != nil {
        response.JSON(c, http.StatusOK, "注册失败", err)
        return
    }

    response.JSON(c, http.StatusOK, "注册成功", nil)
    return
}
