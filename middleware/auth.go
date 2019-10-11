package middleware

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

func AuthSessionMiddle() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userInfo, err := ctx.Cookie("UserInfo")

        if userInfo == "" || err != nil {
            response.JSON(ctx, http.StatusUnauthorized, "未登录，需要登录后操作~", nil)
            ctx.Abort()

            return
        }

        ctx.Next()
    }
}
