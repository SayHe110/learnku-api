package middleware

import (
    "github.com/gin-gonic/gin"
    "learnku-api/pkg/ecode"
    "learnku-api/pkg/jwt"
    "learnku-api/pkg/response"
    "net/http"
    "time"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        code := http.StatusOK
        token := c.Query("token")

        if token == "" {
            code = ecode.InvalidTokenParams.Code
        } else {
            claims, err := jwt.ParseToken(token)
            if err != nil {
                code = ecode.ParseTokenError.Code
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = ecode.TokenExpireTime.Code
            }
        }

        if code != http.StatusOK {
            response.JSON(c, http.StatusOK, ecode.GetECode(code), nil)

            c.Abort()
            return
        }

        c.Next()
    }
}
