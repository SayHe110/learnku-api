package users

import (
    "errors"
    "github.com/gin-gonic/gin"
    usersHandle "learnku-api/handle/users"
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

func Store(c *gin.Context, param *users.UserStoreParam) (err error) {

    rePassword := c.Request.Form.Get("re_password")

    if param.Password != rePassword {
        return errors.New("两次输入的密码不一致哦~")
    }

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

func Login() {

}
