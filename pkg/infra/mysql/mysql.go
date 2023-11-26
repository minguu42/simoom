// Package mysql は MySQL の操作に関するパッケージ
package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/simoom/pkg/config"
)

// Client は repository.Repository を満たす MySQL クライアント
type Client struct {
	db *sql.DB
}

// Close は新しいクエリの実行を停止し、MySQL サーバとの接続を閉じる
func (c *Client) Close() error {
	return c.db.Close()
}

// NewClient は MySQL サーバとの接続を確立し、クライアントを初期化する
func NewClient(conf config.MySQL) (*Client, error) {
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
		if err := db.Ping(); err != nil && i != maxFailureTimes {
			log.Println("db.Ping failed. try again after 15 seconds")
			time.Sleep(15 * time.Second)
			continue
		} else if err != nil && i == maxFailureTimes {
			return nil, fmt.Errorf("failed to verify a connection to the database: %w", err)
		}
		break
	}

	return &Client{db: db}, nil
}
