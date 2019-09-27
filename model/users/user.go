package users

type Users struct {
    // gorm.Model
    ID                int64  `json:"id"`
    Name              string `json:"name"`
    Gender            int    `json:"gender"`
    Mobile            string `json:"mobile"`
    Email             string `json:"email"`
    Password          string `json:"password"`
    RememberToken     string `json:"remember_token"`
    Avatar            string `json:"avatar"`
    Introduction      string `json:"introduction"`
    NotificationCount int    `json:"notification_count"`
    LastActivedAt     string `json:"last_actived_at"`
}
