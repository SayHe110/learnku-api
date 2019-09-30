package topics

import (
    "learnku-api/bootstrap"
)

func GetTopicList() (res []struct{}, err error) {
    res = []struct{}{}

    if err = bootstrap.DB.Self.Table("topics").Select("id, title, category_id, like_count, reply_count, created_at").Order("like_count").Scan(&res).Error;
        err != nil {
        return nil, err
    }

    return res, nil
}
