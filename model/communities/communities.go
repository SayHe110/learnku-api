package communities

import "time"

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
