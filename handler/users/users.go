package users

import (
    "github.com/jinzhu/gorm"
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
    if err = h.db.Limit(2).Offset(0).Find(&res).Error; err != nil {
        log.Printf("select users field (%v)", err)
        return nil, err
    }
    return
}

func (h *Handler) StoreUser(param *userModel.UserStoreParam) (err error) {
    if err = h.db.Exec(_userStoreSQL, param.Email, param.Password,
        time.Now().Format("2006-01-02 15:04:05")).Error; err != nil {
        return
    }
    return
}

func (h *Handler) LoginCheckEmail(param *userModel.UserLoginParam) (res bool, err error) {
    resCount := 0
    if err = h.db.Model(&userModel.Users{}).Where("email = ?", param.Email).Count(&resCount).Error; err != nil {
        return false, err
    }

    if resCount > 0 {
        return true, nil
    }
    return false, nil
}

func (h *Handler) LoginByEmail(param *userModel.UserLoginParam) (res bool, user *userModel.Users, err error) {
    user = &userModel.Users{}

    if err = h.db.Where("email = ?", param.Email).First(&user).Error; err != nil {
        return false, nil, nil
    }

    if ok, _ := auth.DecodePwd(user.Password, param.Password); ok {
        return true, user, nil
    }
    return false, nil, nil
}
