package http

import (
    "github.com/gin-gonic/gin"
    "learnku-api/model/users"
    "learnku-api/pkg/auth"
    "learnku-api/pkg/ecode"
    "learnku-api/pkg/jwt"
    "learnku-api/pkg/response"
    "net/http"
)

func userList(c *gin.Context) {
    res, _ := userSvc.Users()
    response.JSON(c, http.StatusOK, "获取成功", res)
}

func userLogout(c *gin.Context) {
    if ok := auth.CheckAuthLoginStatus(c); !ok {
        response.JSON(c, 50000, "用户未登陆", nil)
        return
    }

    auth.ClearAuthSession(c)
    response.JSON(c, http.StatusNoContent, "退出成功", nil)
}

func userRefreshToken(c *gin.Context) {
    token, err := jwt.RefreshToken(c.Request.Header.Get("token"))
    if err != nil {
        response.JSON(c, ecode.RefreshTokenError.Code, ecode.RefreshTokenError.Message, nil)
    }

    response.JSON(c, http.StatusOK, "刷新token成功", gin.H{
        "token": token,
    })
}

func userStore(c *gin.Context) {
    userRes := &users.UserStoreParam{}

    if err := c.ShouldBind(&userRes); err != nil {
        response.JSON(c, 50001, "参数绑定错误", nil)
        return
    }

    if err := userRes.UserStoreValidator().Errors; err != nil {
        response.JSON(c, 50002, "参数验证错误", userRes.UserStoreValidator())
        return
    }

    if err := userSvc.Store(userRes); err != nil {
        response.JSON(c, 50003, "注册失败", err.Error())
        return
    }

    gerToken, err := jwt.GenerateToken(userRes.Email, userRes.Password)
    if err != nil {
        response.JSON(c, 50004, "生成 token 错误", err.Error())
        return
    }

    response.JSON(c, http.StatusOK, "注册成功", gin.H{
        "token": gerToken,
    })
    return
}

func userLogin(c *gin.Context) {
    userRes := &users.UserLoginParam{}

    if err := c.ShouldBind(&userRes); err != nil {
        response.JSON(c, 50001, "参数绑定错误", nil)
        return
    }

    if err := userRes.UserLoginValidator().Errors; err != nil {
        response.JSON(c, 50002, "参数验证错误", userRes.UserLoginValidator())
        return
    }

    user, err := userSvc.Login(userRes);
    if err != nil {
        response.JSON(c, 50003, err.Error(), nil)
        return
    }

    gerToken, err := jwt.GenerateToken(userRes.Email, userRes.Password)
    if err != nil {
        response.JSON(c, 50004, "生成 token 错误", err.Error())
        return
    }

    // set session
    auth.SaveAuthSession(c, user)

    response.JSON(c, http.StatusOK, "登录成功", gin.H{
        "token": gerToken,
    })
    return
}
