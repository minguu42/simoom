package model

import "time"

type Task struct {
	ID          string
	Steps       []Step
	Tags        []Tag
	UserID      string
	ProjectID   string
	Name        string
	Content     string
	Priority    uint
	DueOn       *time.Time
	CompletedAt *time.Time
}
