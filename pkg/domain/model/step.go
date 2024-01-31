package model

import "time"

type Step struct {
	ID          string
	UserID      string
	TaskID      string
	Name        string
	CompletedAt *time.Time
}
