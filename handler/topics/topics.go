package topics

import (
    "github.com/jinzhu/gorm"
    "learnku-api/config"
    topicsModel "learnku-api/model/topics"
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

func (h *Handler) GetTopicList() (res []*topicsModel.Topics, err error) {
    if err = h.db.Preload("UserInfo").Find(&res).Error; err != nil {
        return nil, err
    }

    return
}

func (h *Handler) GetTopicById(id string) (res []*topicsModel.Topics, err error) {
    if err = h.db.Where("id = ?", id).Preload("UserInfo").Find(&res).Error; err != nil {
        return nil, err
    }

    return
}
