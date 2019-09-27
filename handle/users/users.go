package users

import (
    "learnku-api/bootstrap"
    "learnku-api/model/users"
    "log"
)

func GetUserList() (res *users.Users, err error) {
    res = &users.Users{}

    if err = bootstrap.DB.Self.Find(&res).Error; err != nil {
        log.Printf("select users field (%v)", err)
        return nil, err
    }
    return
}
