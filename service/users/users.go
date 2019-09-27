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

func Store(param *users.UserStoreParam) (err error) {
    if err = usersHandle.StoreUser(param); err != nil {
        return err
    }

    return
}
