package topics

import (
    "learnku-api/config"
    topicsHandler "learnku-api/handler/topics"
    "learnku-api/model/topics"
)

type Service struct {
    handler *topicsHandler.Handler
}

func New(c *config.Config) (s *Service) {
    s = &Service{
        handler: topicsHandler.New(c),
    }

    return
}

func (s *Service) GetTopicsList() (res []*topics.Topics, err error) {
    res, err = s.handler.GetTopicList()
    return
}

func (s *Service) GetTopicsById(id string) (res []*topics.Topics, err error) {
    res, err = s.handler.GetTopicById(id)
    return
}
