package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/api/apperr"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const languageJapanese = "ja-JP"

// NewArrangeError はレスポンス用のエラーを生成するインターセプタを返す
func NewArrangeError() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				return resp, connectError(err)
			}
			return resp, nil
		}
	}
}

func connectError(err error) *connect.Error {
	var (
		code            connect.Code
		message         string
		messageJapanese string
	)
	switch {
	case errors.Is(err, apperr.ErrInvalidAuthorizationFormat):
		code = connect.CodeUnauthenticated
		message = apperr.ErrInvalidAuthorizationFormat.Error()
		messageJapanese = "Authorizationヘッダは'Bearer xxx'の形式で認証トークンを指定してください。"
	case errors.Is(err, apperr.ErrAuthenticationFailed):
		code = connect.CodeUnauthenticated
		message = apperr.ErrAuthenticationFailed.Error()
		messageJapanese = "ユーザの認証に失敗しました。"
	case errors.Is(err, apperr.ErrProjectNotFound):
		code = connect.CodeNotFound
		message = apperr.ErrProjectNotFound.Error()
		messageJapanese = "指定したプロジェクトは見つかりません。"
	case errors.Is(err, apperr.ErrStepNotFound):
		code = connect.CodeNotFound
		message = apperr.ErrStepNotFound.Error()
		messageJapanese = "指定したステップは見つかりません。"
	case errors.Is(err, apperr.ErrTagNotFound):
		code = connect.CodeNotFound
		message = apperr.ErrTagNotFound.Error()
		messageJapanese = "指定したタグは見つかりません。"
	case errors.Is(err, apperr.ErrTaskNotFound):
		code = connect.CodeNotFound
		message = apperr.ErrTaskNotFound.Error()
		messageJapanese = "指定したタスクは見つかりません。"
	case errors.Is(err, apperr.ErrUserNotFound):
		code = connect.CodeNotFound
		message = apperr.ErrUserNotFound.Error()
		messageJapanese = "指定したユーザは見つかりません。"
	default:
		code = connect.CodeUnknown
		message = "some error has occurred on the server side. please wait a few minutes and try again"
		messageJapanese = "サーバ側で何らかのエラーが発生しました。時間を置いてから再度お試しください。"
	}
	connectErr := connect.NewError(code, errors.New(message))
	d, _ := connect.NewErrorDetail(&errdetails.LocalizedMessage{
		Locale:  languageJapanese,
		Message: messageJapanese,
	})
	connectErr.AddDetail(d)
	return connectErr
}
