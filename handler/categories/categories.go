package categories

import (
    "github.com/jinzhu/gorm"
    "learnku-api/bootstrap"
    "learnku-api/config"
    "learnku-api/model/categories"
    "learnku-api/pkg/db"
    "time"
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

func Store(params categories.CategoryParams) (err error) {
    params.CreatedAt = time.Now().Unix()

    if err = bootstrap.DB.Self.Table("categories").Create(&params).Error; err != nil {
        return
    }

    return nil
}
