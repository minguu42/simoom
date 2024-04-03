package model

import "time"

type TaskID string

type Task struct {
	ID          TaskID
	Steps       []Step
	Tags        []Tag
	UserID      UserID
	ProjectID   ProjectID
	Name        string
	Content     string
	Priority    uint
	DueOn       *time.Time
	CompletedAt *time.Time
}
