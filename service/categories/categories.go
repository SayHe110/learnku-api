package categories

import (
    "learnku-api/config"
    categoriesHandle "learnku-api/handler/categories"
    categoryModel "learnku-api/model/categories"
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

func (s *Service) Store(params *categoryModel.CategoryParams) (err error) {
    if err = s.handler.Store(params); err != nil {
        return
    }

    return nil
}

func (s *Service) List() (res []*categoryModel.Categories, err error) {
    res, err = s.handler.List()

    return
}
