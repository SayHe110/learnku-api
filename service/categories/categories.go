package categories

import (
    categoriesHandle "learnku-api/handler/categories"
    "learnku-api/model/categories"
)

func Store(params categories.CategoryParams) (err error) {
    if err = categoriesHandle.Store(params); err != nil {
        return err
    }

    return nil
}
