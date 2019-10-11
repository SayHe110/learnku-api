package communities

import (
    validatorPkg "learnku-api/pkg/validator"
    "time"
)

// 社区的分类 （PHP、MySQL）
type CommCategories struct {
    Id             int           `json:"id"`
    Name           string        `json:"name"`
    Description    string        `json:"description"`
    CreatedAt      int           `json:"created_at"`
    UpdatedAt      time.Time     `json:"updated_at"`
    DeletedAt      time.Time     `json:"-"`
    CommunitiesRes []Communities `json:"communities_res" gorm:"foreignkey:CommCategoriesId"`
}

type Communities struct {
    Id               int       `json:"id"`
    CommCategoriesId int       `json:"comm_categories_id"`
    Name             string    `json:"name"`
    Description      string    `json:"description"`
    Img              string    `json:"img"`
    CreatedAt        int       `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
    DeletedAt        time.Time `json:"-"`
}

type CommunityParams struct {
    Id          int    `form:"id" binding:"required" validate:"is_filed_exist"`
    Name        string `form:"name" binding:"required" validate:"max=25,min=2"`
    Description string `form:"description" binding:"required"`
}

func (c *CommunityParams) UpdateCommunityValidator() validatorPkg.CommonError {
    return validatorPkg.CustomValParams(c, validatorPkg.FiledValExists, "is_filed_exist", "communities")
}
