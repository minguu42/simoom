package model

import "time"

type Step struct {
	ID          string
	UserID      string
	TaskID      string
	Title       string
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
