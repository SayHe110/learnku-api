package middleware

import (
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/response"
    "net/http"
)

//TODO 存入后获取不到，在存入时出错（转换为 string 时）
func AuthSessionMiddle() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        session := sessions.Default(ctx)
        sessionValue := session.Get("UserInfo").([]byte)

        if sessionValue == nil {
            response.JSON(ctx, http.StatusUnauthorized, "未登录，需要登录后操作~", nil)

            ctx.Abort()
            return
        }

        // ctx.Set("UserInfo", json.Unmarshal(sessionValue, &users.Users{}))
        // res := session.Get("UserInfo")
        //user := &users.Users{}
        //_ = json.Unmarshal(sessionValue, user)
        //log.Println(user)
        ctx.Next()
    }
}
