package topics

import (
    "learnku-api/bootstrap"
    "learnku-api/model/topics"
)

func GetTopicList() (res []*topics.Topics, err error) {
    if err = bootstrap.DB.Self.Preload("UserInfo").Find(&res).Error; err != nil {
        return nil, err
    }

    return
}

func GetTopicById(id string) (res []*topics.Topics, err error) {
    if err = bootstrap.DB.Self.Where("id = ?", id).Preload("UserInfo").Find(&res).Error; err != nil {
        return nil, err
    }

    return
}
