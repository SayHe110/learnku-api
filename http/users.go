package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/model/users"
    "learnku-api/pkg/response"
    userSvc "learnku-api/service/users"
    "learnku-api/tools"
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

    res := tools.ValidatorCommonError{}
    res = userRes.UserStoreValidator()
    if res.Errors != nil {
        response.JSON(c, 50002, "参数验证错误" ,userRes.UserStoreValidator())
        return
    }

    if err := userSvc.Store(c, userRes); err != nil {
        response.JSON(c, 50003, "注册失败", err.Error())
        return
    }

    response.JSON(c, http.StatusOK, "注册成功", nil)
    return
}

func userLogin(c *gin.Context) {
    userSvc.Login()
}
