package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/minguu42/simoom/api/logging"
	proxy "github.com/shogo82148/go-sql-proxy"
)

func registerDriverWithHooks(name string) {
	sql.Register(name, proxy.NewProxyContext(&mysql.MySQLDriver{}, &proxy.HooksContext{
		PreExec: func(_ context.Context, _ *proxy.Stmt, _ []driver.NamedValue) (any, error) {
			return time.Now(), nil
		},
		PostExec: func(ctx context.Context, v any, stmt *proxy.Stmt, args []driver.NamedValue, _ driver.Result, err error) error {
			if errors.Is(err, driver.ErrSkip) {
				return nil
			}

			values := make([]any, 0, len(args))
			for _, arg := range args {
				values = append(values, arg.Value)
			}
			logging.SQL(ctx, stmt.QueryString, values, time.Since(v.(time.Time)), err)
			return nil
		},
		PreQuery: func(_ context.Context, _ *proxy.Stmt, _ []driver.NamedValue) (any, error) {
			return time.Now(), nil
		},
		PostQuery: func(ctx context.Context, v any, stmt *proxy.Stmt, args []driver.NamedValue, _ driver.Rows, err error) error {
			if errors.Is(err, driver.ErrSkip) {
				return nil
			}

			values := make([]any, 0, len(args))
			for _, arg := range args {
				values = append(values, arg.Value)
			}
			logging.SQL(ctx, stmt.QueryString, values, time.Since(v.(time.Time)), err)
			return nil
		},
	}))
}
