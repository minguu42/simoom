package model

type TagID string

type Tag struct {
	ID     TagID
	UserID UserID
	Name   string
}
