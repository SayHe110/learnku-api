package users

import "time"

type Users struct {
    ID                int64     `json:"id"`
    Name              string    `json:"name"`
    Gender            int       `json:"gender"`
    Mobile            string    `json:"mobile"`
    Email             string    `json:"email"`
    Password          string    `json:"-"`
    RememberToken     string    `json:"remember_token"`
    Avatar            string    `json:"avatar"`
    Introduction      string    `json:"introduction"`
    NotificationCount int       `json:"notification_count"`
    LastActivedAt     string    `json:"last_actived_at"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
    DeletedAt         time.Time `json:"deleted_at"`
}
