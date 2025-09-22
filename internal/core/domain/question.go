package domain

import "time"

type Question struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Content   string    `json:"content" gorm:"not null"`
    Answer    string    `json:"answer"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}