package middleware

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "learnku-api/model/users"
    casbinPkg "learnku-api/pkg/casbin"
    "learnku-api/pkg/response"
    "net/http"
)

func AuthSessionMiddle(obj, act string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userInfo, err := ctx.Cookie("UserInfo")

        user := users.Users{}
        _ = json.Unmarshal([]byte(userInfo), &user)

        if userInfo == "" || err != nil {
            response.JSON(ctx, http.StatusUnauthorized, "未登录，需要登录后操作~", nil)
            ctx.Abort()

            return
        }

        ok, err := casbinPkg.Enforce(user.Email, obj, act)
        if err != nil {
            response.JSON(ctx, http.StatusInternalServerError, err.Error(), nil)
            ctx.Abort()

            return
        }

        if ! ok {
            response.JSON(ctx, http.StatusForbidden, "用户没有权限", nil)
            ctx.Abort()

            return
        }

        ctx.Next()
    }
}
