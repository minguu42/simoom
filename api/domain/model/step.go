package model

import "time"

type StepID string

type Step struct {
	ID          StepID
	UserID      UserID
	TaskID      TaskID
	Name        string
	CompletedAt *time.Time
}
