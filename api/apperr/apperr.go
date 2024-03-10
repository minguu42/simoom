// Package apperr はアプリケーション独自のエラーを定義する
// このパッケージで定義されたエラーは handler パッケージと usecase パッケージから利用する
package apperr

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type Error struct {
	err             error
	code            connect.Code
	message         string
	messageJapanese string
}

func (e Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.message, e.err)
	}
	return e.message
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
		code:            connect.CodeInvalidArgument,
		message:         "the entered value is incorrect",
		messageJapanese: "入力された値に誤りを含みます",
	}
}

func ErrInvalidAuthorizationFormat() Error {
	return Error{
		code:            connect.CodeUnauthenticated,
		message:         "the Authorization header should include a value in the form 'Bearer xxx'",
		messageJapanese: "Authorizationヘッダには'Bearer xxx'の形式で認証トークンを指定してください",
	}
}

func ErrAuthentication(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeUnauthenticated,
		message:         "authentication failed",
		messageJapanese: "ユーザの認証に失敗しました",
	}
}

func ErrProjectNotFound(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeNotFound,
		message:         "the specified project is not found",
		messageJapanese: "指定したプロジェクトは見つかりません",
	}
}
func ErrStepNotFound(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeNotFound,
		message:         "the specified step is not found",
		messageJapanese: "指定したステップは見つかりません",
	}
}
func ErrTagNotFound(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeNotFound,
		message:         "the specified tag is not found",
		messageJapanese: "指定したタグは見つかりません",
	}
}
func ErrTaskNotFound(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeNotFound,
		message:         "the specified task is not found",
		messageJapanese: "指定したタスクは見つかりません",
	}
}

func ErrUserNotFound(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeNotFound,
		message:         "the specified user is not found",
		messageJapanese: "指定したユーザは見つかりません",
	}
}

func ErrUnknown(err error) Error {
	return Error{
		err:             err,
		code:            connect.CodeUnknown,
		message:         "some error has occurred on the server side. please wait a few minutes and try again",
		messageJapanese: "サーバ側で何らかのエラーが発生しました。時間を置いてから再度お試しください。",
	}
}
