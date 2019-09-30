package topics

import topicsSvc "learnku-api/handler/topics"

func Topics() (res []struct{}, err error) {
    res, err = topicsSvc.GetTopicList()
    return
}
