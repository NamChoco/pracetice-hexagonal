package domain

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    GoogleID  string    `json:"google_id" gorm:"uniqueIndex"`
    Email     string    `json:"email" gorm:"uniqueIndex"`
    Name      string    `json:"name"`
    Picture   string    `json:"picture"`
    Role      UserRole  `json:"role" gorm:"default:user"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserRole string

const (
    UserRoleUser  UserRole = "user"
    UserRoleAdmin UserRole = "admin"
)

type UserSession struct {
    UserID    uint      `json:"user_id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Role      UserRole  `json:"role"`
    ExpiresAt time.Time `json:"expires_at"`
}

// OAuth User representation from Google
type GoogleUser struct {
    GoogleID string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Picture  string `json:"picture"`
}