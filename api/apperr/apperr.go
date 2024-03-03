// Package apperr はアプリケーション独自のエラーを定義する
// このパッケージで定義されたエラーは handler パッケージと usecase パッケージから利用する
package apperr

import "errors"

var (
	ErrInvalidAuthorizationFormat = errors.New("the Authorization header should include a value in the form 'Bearer xxx'")
	ErrAuthenticationFailed       = errors.New("authentication failed")

	ErrProjectNotFound = errors.New("the specified project is not found")
	ErrStepNotFound    = errors.New("the specified step is not found")
	ErrTagNotFound     = errors.New("the specified tag is not found")
	ErrTaskNotFound    = errors.New("the specified task is not found")
	ErrUserNotFound    = errors.New("the specified user is not found")
)
