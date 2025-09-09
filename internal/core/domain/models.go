package domain

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Provider   string
	ProviderID string `gorm:"index"`
	Email      string `gorm:"uniqueIndex"`
	Name       string
	IsAdmin    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Question struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Body         string
	Answer       *string
	AskedByID    uint
	AskedBy      *User `gorm:"foreignKey:AskedByID"`
	AnsweredByID *uint
	AnsweredBy   *User `gorm:"foreignKey:AnsweredByID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
