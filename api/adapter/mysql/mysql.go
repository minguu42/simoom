package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/minguu42/simoom/api/adapter/mysql/sqlc"
	"github.com/minguu42/simoom/api/config"
)

type transactionKey struct{}

// Client は repository.Repository を満たすMySQLクライアント
type Client struct {
	db *sql.DB
}

// Close は新しいクエリの実行を停止し、MySQLサーバとの接続を閉じる
func (c *Client) Close() error {
	return c.db.Close()
}

// NewClient はMySQLサーバとの接続を確立したクライアントを返す
func NewClient(conf config.DB) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	name := "mysql-proxy"
	registerDriverWithHooks(name)
	db, err := sql.Open(name, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database: %w", err)
	}
	db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeMin) * time.Minute)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	maxRetryCount := 20
	for c := range maxRetryCount {
		if err := db.Ping(); err == nil {
			break
		} else if c == maxRetryCount-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		time.Sleep(1 * time.Second)
	}
	return &Client{db: db}, nil
}

// queries は ctx から *sqlc.Queries を取得する
// ctx に *sqlc.Queries が存在しない場合は新しく生成し、返す
func (c *Client) queries(ctx context.Context) *sqlc.Queries {
	q, ok := ctx.Value(transactionKey{}).(*sqlc.Queries)
	if ok {
		return q
	}
	return sqlc.New(c.db)
}

// Transaction は fn を1つのトランザクション内で実行する
func (c *Client) Transaction(ctx context.Context, fn func(ctxWithTx context.Context) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	ctxWithTxQueries := context.WithValue(ctx, transactionKey{}, sqlc.New(tx))

	if err := fn(ctxWithTxQueries); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("error occurred during the transaction: %w, and rollback failed: %w", err, rollbackErr)
		}
		return fmt.Errorf("error occurred during the transaction and it was rolled back: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
