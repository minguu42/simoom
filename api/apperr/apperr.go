package apperr

import (
	"context"
	"errors"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func NewError(err error) Error {
	var appErr Error
	switch {
	case errors.As(err, &appErr):
		return appErr
	case errors.Is(err, context.DeadlineExceeded), errors.Is(err, context.Canceled):
		return ErrDeadlineExceeded(err)
	default:
		return ErrUnknown(err)
	}
}

type Error struct {
	err             error
	id              string
	code            connect.Code
	message         string
	messageJapanese string
}

func (e Error) IsZero() bool {
	return e.id == ""
}

func (e Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.message, e.err)
	}
	return e.message
}

func (e Error) ID() string {
	return e.id
}

func (e Error) ConnectError() *connect.Error {
	err := connect.NewError(e.code, errors.New(e.message))
	d, _ := connect.NewErrorDetail(&errdetails.LocalizedMessage{
		Locale:  "ja-JP",
		Message: e.messageJapanese,
	})
	err.AddDetail(d)
	return err
}

func ErrInvalidRequest(err error) Error {
	return Error{
		err:             err,
		id:              "invalid-request",
		code:            connect.CodeInvalidArgument,
		message:         "the entered value is incorrect",
		messageJapanese: "入力された値に誤りを含みます",
	}
}

func ErrDuplicateUserName(err error) Error {
	return Error{
		err:             err,
		id:              "duplicate-user-name",
		code:            connect.CodeInvalidArgument,
		message:         "the name is already in use",
		messageJapanese: "そのユーザ名は既に使用されています",
	}
}

func ErrDuplicateUserEmail(err error) Error {
	return Error{
		err:             err,
		id:              "duplicate-user-email",
		code:            connect.CodeInvalidArgument,
		message:         "the mail address is already in use",
		messageJapanese: "そのメールアドレスは既に使用されています",
	}
}

func ErrDeadlineExceeded(err error) Error {
	return Error{
		err:             err,
		id:              "deadline-exceeded",
		code:            connect.CodeDeadlineExceeded,
		message:         "request was not processed within the specified time",
		messageJapanese: "リクエストは規定時間内に処理されませんでした",
	}
}

func ErrInvalidAuthorizationFormat() Error {
	return Error{
		id:              "invalid-authorization-format",
		code:            connect.CodeUnauthenticated,
		message:         "the Authorization header should include a value in the form 'Bearer xxx'",
		messageJapanese: "Authorizationヘッダには'Bearer xxx'の形式で認証トークンを指定してください",
	}
}

func ErrAuthentication(err error) Error {
	t := time.Time{}
	t.IsZero()
	return Error{
		err:             err,
		id:              "authentication",
		code:            connect.CodeUnauthenticated,
		message:         "authentication failed",
		messageJapanese: "ユーザの認証に失敗しました",
	}
}

func ErrProjectNotFound(err error) Error {
	return Error{
		err:             err,
		id:              "project-not-found",
		code:            connect.CodeNotFound,
		message:         "the specified project is not found",
		messageJapanese: "指定したプロジェクトは見つかりません",
	}
}

func ErrStepNotFound(err error) Error {
	return Error{
		err:             err,
		id:              "step-not-found",
		code:            connect.CodeNotFound,
		message:         "the specified step is not found",
		messageJapanese: "指定したステップは見つかりません",
	}
}

func ErrTagNotFound(err error) Error {
	return Error{
		err:             err,
		id:              "tag-not-found",
		code:            connect.CodeNotFound,
		message:         "the specified tag is not found",
		messageJapanese: "指定したタグは見つかりません",
	}
}

func ErrTaskNotFound(err error) Error {
	return Error{
		err:             err,
		id:              "task-not-found",
		code:            connect.CodeNotFound,
		message:         "the specified task is not found",
		messageJapanese: "指定したタスクは見つかりません",
	}
}

func ErrUserNotFound(err error) Error {
	return Error{
		err:             err,
		id:              "user-not-found",
		code:            connect.CodeNotFound,
		message:         "the specified user is not found",
		messageJapanese: "指定したユーザは見つかりません",
	}
}

func ErrUnknown(err error) Error {
	return Error{
		err:             err,
		id:              "unknown",
		code:            connect.CodeUnknown,
		message:         "some error has occurred on the server side. please wait a few minutes and try again",
		messageJapanese: "サーバ側で何らかのエラーが発生しました。時間を置いてから再度お試しください。",
	}
}
