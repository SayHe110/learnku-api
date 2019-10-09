package users

import (
    "errors"
    usersHandle "learnku-api/handler/users"
    "learnku-api/model/users"
    "learnku-api/pkg/auth"
)

func Users() (res []*users.Users, err error) {
    res, err = usersHandle.GetUserList()
    if err != nil {
        return nil, err
    }

    return
}

func Store(param *users.UserStoreParam) (err error) {
    hashPwd, err := auth.EncodePwd(param.Password)
    if err != nil {
        return err
    }

    storeParams := &users.UserStoreParam{
        Email:    param.Email,
        Password: hashPwd,
    }

    if err = usersHandle.StoreUser(storeParams); err != nil {
        return err
    }

    return
}

func Login(param *users.UserLoginParam) (*users.Users, error) {
    if ok, _ := usersHandle.LoginCheckEmail(param); !ok {
        return nil, errors.New("该邮箱不存在或者不正确")
    }

    ok, user, _ := usersHandle.LoginByEmail(param);
    if !ok {
        return nil, errors.New("密码不正确")
    }

    return user, nil
}
