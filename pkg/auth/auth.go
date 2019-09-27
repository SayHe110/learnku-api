package auth

import (
    "golang.org/x/crypto/bcrypt"
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

func DecodePwd(pwd string) (bool, error) {
    return true, nil
}
