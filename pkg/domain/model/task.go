package model

import "time"

type Task struct {
	ID          string
	Steps       []Step
	Tags        []Tag
	ProjectID   string
	Title       string
	Content     string
	Priority    uint
	DueOn       *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
