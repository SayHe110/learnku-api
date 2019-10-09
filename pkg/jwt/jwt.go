package jwt

import (
    jwtGo "github.com/dgrijalva/jwt-go"
    "learnku-api/config"
    "time"
)

type Claims struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    jwtGo.StandardClaims
}

func GenerateToken(email, password string) (tokenString string, err error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(config.C.AppConfig.ExpireTime)

    claims := Claims{
        Email:    email,
        Password: password,
        StandardClaims: jwtGo.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    config.C.AppConfig.Name,
        },
    }

    tokenClaims := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString([]byte(config.C.AppConfig.JWTSecret))

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwtGo.ParseWithClaims(token, &Claims{}, func(token *jwtGo.Token) (interface{}, error) {
        return []byte(config.C.AppConfig.JWTSecret), nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}

//TODO 该方法刷新方式为重新创建 token，应该是将之前的 token 增加过期时间，而不是重新给 token，查看文档并没有解决问题
func RefreshToken(token string) (refreshToken string, err error) {
    tokenClaims, err := jwtGo.ParseWithClaims(token, &Claims{}, func(token *jwtGo.Token) (interface{}, error) {
        return []byte(config.C.AppConfig.JWTSecret), nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            refreshToken, err = GenerateToken(claims.Email, claims.Email)

            return
        }
    }

    return "", err
}
