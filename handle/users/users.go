package users

import (
    "learnku-api/bootstrap"
    userModel "learnku-api/model/users"
    "learnku-api/pkg/auth"
    "log"
)

func GetUserList() (res []*userModel.Users, err error) {
    if err = bootstrap.DB.Self.Limit(2).Offset(0).Find(&res).Error; err != nil {
        log.Printf("select users field (%v)", err)
        return nil, err
    }
    return
}

func StoreUser(param *userModel.UserStoreParam) (err error) {

    hashPwd, err := auth.EncodePwd(param.Password)
    if err != nil {
        return err
    }

    userTmp := &userModel.UserStoreParam{
        Email:    param.Email,
        Password: hashPwd,
    }

    if err = bootstrap.DB.Self.Table("users").Create(&userTmp).Error; err != nil {
        log.Printf("store user field (%v)", err)
        return err
    }

    return
}
