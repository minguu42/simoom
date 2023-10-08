// Package usecase はユースケースを定義する
package usecase

import "github.com/cockroachdb/errors"

const userID = "01DXF6DT000000000000000000"

var (
	ErrProjectNotFound = errors.New("the specified project is not found")
	ErrStepNotFound    = errors.New("the specified step is not found")
	ErrTagNotFound     = errors.New("the specified tag is not found")
	ErrTaskNotFound    = errors.New("the specified task is not found")
	ErrUnkown          = errors.New("an unintentional error occurred on the server")
)
