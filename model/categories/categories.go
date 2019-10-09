package categories

import (
    validatorPkg "learnku-api/pkg/validator"
    "time"
)

type Categories struct {
    Id          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   string    `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    DeletedAt   time.Time `json:"-"`
}

type CategoryParams struct {
    Name        string `form:"name" binding:"required" validate:"max=25,min=10,filed_unique"`
    Description string `form:"description" binding:"required"`
    CreatedAt   int64
}

func (params *CategoryParams) CategoryStoreValidator() (valError validatorPkg.CommonError) {
    return validatorPkg.CustomValParams(params, validatorPkg.FiledValUnique, "filed_unique", "categories")
}
