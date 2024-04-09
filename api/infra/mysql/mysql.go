// Package mysql はMySQLを扱うパッケージ
package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/simoom/api/config"
	"github.com/minguu42/simoom/api/infra/mysql/sqlc"
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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database: %w", err)
	}
	db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeMin) * time.Minute)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	maxFailureTimes := 2
	for i := 0; i <= maxFailureTimes; i++ {
		if err := db.Ping(); err != nil {
			if i == maxFailureTimes {
				return nil, fmt.Errorf("failed to verify a connection to the database: %w", err)
			}
			log.Println("db.Ping failed. try again after 15 seconds")
			time.Sleep(15 * time.Second)
			continue
		}
		break
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
