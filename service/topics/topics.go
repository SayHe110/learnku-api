package topics

import (
    topicsSvc "learnku-api/handler/topics"
    "learnku-api/model/topics"
)

func Topics() (res []*topics.Topics, err error) {
    res, err = topicsSvc.GetTopicList()
    return
}
