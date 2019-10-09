package categories

import (
    "learnku-api/config"
    categoriesHandle "learnku-api/handler/categories"
    "learnku-api/model/categories"
)

type Service struct {
    handler *categoriesHandle.Handler
}

func New(c *config.Config) (s *Service) {
    s = &Service{
        handler: categoriesHandle.New(c),
    }

    return
}

func (s *Service) Store(params categories.CategoryParams) (err error) {
    if err = categoriesHandle.Store(params); err != nil {
        return err
    }

    return nil
}
