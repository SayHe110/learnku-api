package topics

import (
    "learnku-api/bootstrap"
    "learnku-api/model/topics"
)

func GetTopicList() (res []*topics.Topics, err error) {
    // res = topics.Topics{}

    //if err = bootstrap.DB.Self.Table("topics").Select("id, title, category_id, like_count, reply_count, created_at").Order("like_count").Scan(&res).Error;
    //    err != nil {
    //    return nil, err
    //}

    if err = bootstrap.DB.Self.Find(&res).Error; err != nil {
        return nil, err
    }

    return
}
