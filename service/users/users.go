package users

import (
    usersHandle "learnku-api/handle/users"
    "learnku-api/model/users"
)

func Users() (res []*users.Users, err error) {
    res, err = usersHandle.GetUserList()
    if err != nil {
        return nil, err
    }

    return
}

func Store() (err error) {
    if err = usersHandle.StoreUser(); err != nil {
        return err
    }

    return
}