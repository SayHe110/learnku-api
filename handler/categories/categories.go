package categories

import (
    "learnku-api/bootstrap"
    "learnku-api/model/categories"
    "time"
)

func Store(params categories.CategoryParams) (err error) {
    params.CreatedAt = time.Now().Unix()

    if err = bootstrap.DB.Self.Table("categories").Create(&params).Error; err != nil {
        return
    }

    return nil
}
