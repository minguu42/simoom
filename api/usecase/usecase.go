// Package usecase はユースケースを定義する
package usecase

import "errors"

var (
	ErrProjectNotFound = errors.New("the specified project is not found")
	ErrStepNotFound    = errors.New("the specified step is not found")
	ErrTagNotFound     = errors.New("the specified tag is not found")
	ErrTaskNotFound    = errors.New("the specified task is not found")
)
