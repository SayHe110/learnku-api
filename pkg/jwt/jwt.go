package jwt

import (
    jwtGo "github.com/dgrijalva/jwt-go"
    "learnku-api/config"
    "time"
)

// var jwtSecret = []byte(config.AppConfig.JWTSecret)

type Claims struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    // StandardClaims *jwtGo.StandardClaims
    jwtGo.StandardClaims
}

func GenerateToken(email, password string) (tokenString string, err error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(config.AppConfig.ExpireTime)
    //
    //token := jwtGo.New(jwtGo.SigningMethodHS256)
    //claims := make(jwtGo.MapClaims)
    //claims["exp"] = expireTime
    //claims["iat"] = time.Now().Unix()
    //token.Claims = claims
    //
    //if tokenString, err = token.SignedString([]byte(config.AppConfig.JWTSecret)); err != nil {
    //    return "", err
    //}
    //return tokenString, nil

    claims := Claims{
        Email:    email,
        Password: password,
        StandardClaims: jwtGo.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    config.AppConfig.Name,
        },
    }

    tokenClaims := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString([]byte(config.AppConfig.JWTSecret))

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwtGo.ParseWithClaims(token, &Claims{}, func(token *jwtGo.Token) (interface{}, error) {
        return []byte(config.AppConfig.JWTSecret), nil
    })
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}
