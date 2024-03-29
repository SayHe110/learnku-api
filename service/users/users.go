package users

import (
    "errors"
    "github.com/gin-gonic/gin"
    "learnku-api/config"
    usersHandle "learnku-api/handler/users"
    "learnku-api/model/users"
    "learnku-api/pkg/auth"
)

type Service struct {
    handler *usersHandle.Handler
}

func New(c *config.Config) (s *Service) {
    s = &Service{
        handler: usersHandle.New(c),
    }

    return
}

func (s *Service) Users() (res []*users.Users, err error) {
    res, err = s.handler.GetUserList()
    if err != nil {
        return nil, err
    }

    return
}

func (s *Service) Store(param *users.UserStoreParam) (err error) {
    hashPwd, err := auth.EncodePwd(param.Password)
    if err != nil {
        return err
    }

    storeParams := &users.UserStoreParam{
        Email:    param.Email,
        Password: hashPwd,
    }

    if err = s.handler.StoreUser(storeParams); err != nil {
        return err
    }

    return
}

func (s *Service) Login(param *users.UserLoginParam) (*users.Users, error) {
    if ok, _ := s.handler.LoginCheckEmail(param); !ok {
        return nil, errors.New("该邮箱不存在或者不正确")
    }

    ok, user, _ := s.handler.LoginByEmail(param);
    if !ok {
        return nil, errors.New("密码不正确")
    }

    return user, nil
}

func (s *Service) CheckLogin(ctx *gin.Context) bool {
    userInfo, err := ctx.Cookie("UserInfo")
    if userInfo == "" || err != nil {
        return false
    }

    return true
}
