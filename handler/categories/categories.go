package categories

import (
    "github.com/jinzhu/gorm"
    "learnku-api/config"
    categoryModel "learnku-api/model/categories"
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

func (h *Handler) Store(params *categoryModel.CategoryParams) (err error) {
    params.CreatedAt = time.Now().Unix()

    if err = h.db.Table("categories").Create(&params).Error; err != nil {
        return
    }

    return nil
}

func (h *Handler) List() (res []*categoryModel.Categories, err error) {
    if err = h.db.Find(&res).Error; err != nil {
        return nil, err
    }

    return
}
