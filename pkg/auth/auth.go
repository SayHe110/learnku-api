package auth

import (
    "encoding/json"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    userModel "learnku-api/model/users"
    "log"
)

func EncodePwd(pwd string) (string, error) {
    hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

    if err != nil {
        log.Printf("encode password err (%v)", err)
        return "", err
    }

    return string(hashPwd), nil
}

func DecodePwd(hashPwd, pwd string) (bool, error) {
    if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)); err != nil {
        return false, nil
    }

    return true, nil
}

func SaveAuthSession(ctx *gin.Context, users *userModel.Users) {
    session := sessions.Default(ctx)
    seVal, _ := json.Marshal(users)

    session.Set("UserInfo", seVal) // TODO 是否设置过期时间
    _ = session.Save()
}
