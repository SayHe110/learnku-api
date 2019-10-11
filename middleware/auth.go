package middleware

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

// TODO 稍后修改为 Gin 的 cookie
func AuthSessionMiddle() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        session := sessions.Default(ctx)
        sessionValue := session.Get("UserInfo").([]byte)

        if sessionValue == nil {
            response.JSON(ctx, http.StatusUnauthorized, "未登录，需要登录后操作~", nil)

            ctx.Abort()
            return
        }

        ctx.Next()
    }
}
