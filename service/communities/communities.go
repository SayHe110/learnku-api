package communities

import (
    "learnku-api/config"
    communitiesHandle "learnku-api/handler/communities"
    commModel "learnku-api/model/communities"
)

type Service struct {
    handler *communitiesHandle.Handler
}

func New(c *config.Config) (s *Service) {
    s = &Service{
        handler: communitiesHandle.New(c),
    }

    return
}

func (s *Service) GetList() (res []*commModel.CommCategories, err error) {
    if res, err = s.handler.GetList(); err !=nil {
        return nil, err
    }

    return
}
