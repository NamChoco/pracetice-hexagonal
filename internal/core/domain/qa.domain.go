package domain

import "time"

type Question struct {
	ID        uint
	Content   string
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
