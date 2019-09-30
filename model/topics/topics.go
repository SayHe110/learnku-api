package topics

import "time"

type Topics struct {
    ID              uint      `json:"id"`
    Title           string    `json:"title"`
    Body            string    `json:"body"`
    IsEssence       uint      `json:"is_essence"`
    UserId          uint      `json:"user_id"`
    CategoryId      uint      `json:"category_id"`
    ReplyCount      uint      `json:"reply_count"`
    ViewCount       uint      `json:"view_count"`
    LikeCount       uint      `json:"like_count"`
    LastReplyUserId uint      `json:"last_reply_user_id"`
    Order           string    `json:"order"`
    Excerpt         string    `json:"excerpt"`
    Slug            string    `json:"slug"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
    DeletedAt       time.Time `json:"deleted_at"`
}
