package users

import (
    "learnku-api/bootstrap"
    userModel "learnku-api/model/users"
    "log"
)

func GetUserList() (res []*userModel.Users, err error) {

    if err = bootstrap.DB.Self.Limit(2).Offset(0).Find(&res).Error; err != nil {
        log.Printf("select users field (%v)", err)
        return nil, err
    }
    return
}
