package topics

import (
    topicsSvc "learnku-api/handler/topics"
    "learnku-api/model/topics"
)

func Topics() (res []*topics.Topics, err error) {
    res, err = topicsSvc.GetTopicList()
    return
}

func TopicsById(id string) (res []*topics.Topics, err error) {
    res, err = topicsSvc.GetTopicById(id)
    return
}
