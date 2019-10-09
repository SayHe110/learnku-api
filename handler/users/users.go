package users

import (
    "github.com/jinzhu/gorm"
    "learnku-api/bootstrap"
    "learnku-api/config"
    userModel "learnku-api/model/users"
    "learnku-api/pkg/auth"
    "learnku-api/pkg/db"
    "log"
    "time"
)

const (
    _userStoreSQL = `INSERT INTO users(email, password, created_at) VALUES(?,?,?)`
)

type Handler struct {
    db *gorm.DB
}

func New(c *config.Config) (h *Handler) {
    h = &Handler{
        db: db.NewMysql(c),
    }

    return
}

func (h *Handler) GetUserList() (res []*userModel.Users, err error) {
    if err = bootstrap.DB.Self.Limit(2).Offset(0).Find(&res).Error; err != nil {
        log.Printf("select users field (%v)", err)
        return nil, err
    }
    return
}

func StoreUser(param *userModel.UserStoreParam) (err error) {
    if err = bootstrap.DB.Self.Exec(_userStoreSQL, param.Email, param.Password,
        time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
        return
    }
    return
}

func LoginCheckEmail(param *userModel.UserLoginParam) (res bool, err error) {
    resCount := 0
    if err = bootstrap.DB.Self.Model(&userModel.Users{}).Where("email = ?", param.Email).Count(&resCount).Error; err != nil {
        return false, err
    }

    if resCount > 0 {
        return true, nil
    }
    return false, nil
}

func LoginByEmail(param *userModel.UserLoginParam) (res bool, user *userModel.Users, err error) {
    user = &userModel.Users{}

    if err = bootstrap.DB.Self.Where("email = ?", param.Email).First(&user).Error; err != nil {
        return false, nil, nil
    }

    if ok, _ := auth.DecodePwd(user.Password, param.Password); ok {
        return true, user, nil
    }
    return false, nil, nil
}
