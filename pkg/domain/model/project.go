package model

import "time"

type Project struct {
	ID         string
	UserID     string
	Name       string
	Color      string
	IsArchived bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
