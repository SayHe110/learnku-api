package auth

import (
    "encoding/json"
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
    userStr, _ := json.Marshal(users)

    ctx.SetCookie("UserInfo", string(userStr), 3600, "/", "localhost", false, true)
}

func ClearAuthSession(ctx *gin.Context) {
    ctx.SetCookie("UserInfo", "", -1, "/", "localhost", false, true)
}

func CheckAuthLoginStatus(ctx *gin.Context) bool {
    if userInfo, err := ctx.Cookie("UserInfo"); userInfo == "" || err != nil {
        return false
    }

    return true
}
