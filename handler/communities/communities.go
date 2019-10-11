package communities

import (
    "github.com/jinzhu/gorm"
    "learnku-api/config"
    commModel "learnku-api/model/communities"
    "learnku-api/pkg/db"
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

func (h *Handler) GetList() (res []*commModel.CommCategories, err error) {
    if err = h.db.Preload("CommunitiesRes").Find(&res).Error; err != nil {
        return nil, err
    }

    return
}
