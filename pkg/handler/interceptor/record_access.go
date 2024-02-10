package interceptor

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/minguu42/simoom/pkg/applog"
	"github.com/minguu42/simoom/pkg/usecase"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const languageJapanese = "ja-JP"

// NewArrangeErrorAndRecordAccess はリクエスト毎のアクセスログ/エラーログを出力するインターセプタを返す
func NewArrangeErrorAndRecordAccess() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			if err != nil {
				connectErr := connectError(err)
				applog.LogAccessError(ctx, connectErr.Code(), req.Spec().Procedure, err)
				return resp, connectErr
			}
			applog.LogAccess(ctx, req.Spec().Procedure)
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
	case errors.Is(err, usecase.ErrProjectNotFound):
		code = connect.CodeNotFound
		message = usecase.ErrProjectNotFound.Error()
		messageJapanese = "指定したプロジェクトは見つかりません"
	case errors.Is(err, usecase.ErrStepNotFound):
		code = connect.CodeNotFound
		message = usecase.ErrStepNotFound.Error()
		messageJapanese = "指定したステップは見つかりません"
	case errors.Is(err, usecase.ErrTagNotFound):
		code = connect.CodeNotFound
		message = usecase.ErrTagNotFound.Error()
		messageJapanese = "指定したタグは見つかりません"
	case errors.Is(err, usecase.ErrTaskNotFound):
		code = connect.CodeNotFound
		message = usecase.ErrTaskNotFound.Error()
		messageJapanese = "指定したタスクは見つかりません"
	default:
		code = connect.CodeUnknown
		message = "some error has occurred on the server side. please wait a few minutes and try again"
		messageJapanese = "サーバ側で何らかのエラーが発生しました。時間を置いてから再度お試しください。"
	}
	connectErr := connect.NewError(code, errors.New(message))
	d, err := connect.NewErrorDetail(&errdetails.LocalizedMessage{
		Locale:  languageJapanese,
		Message: messageJapanese,
	})
	connectErr.AddDetail(d)
	return connectErr
}
