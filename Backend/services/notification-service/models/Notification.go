package models

type Notification struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    UserID    uint   `json:"user_id"`
    Message   string `json:"message"`
    Read      bool   `json:"read"`
    CreatedAt string `json:"created_at"`
}