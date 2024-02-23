package mysql

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-testfixtures/testfixtures/v3"
)

// Migrate はinfra/mysql/schema.sqlを読み込んで、データベースのマイグレーションを行う
// エラー時はログを出力し、os.Exit(1)で終了する
func Migrate(client *Client) {
	// os.Getwd ではテスト時にパッケージ毎に得られるファイルパスが動的に変化するため、runtime.Caller を使用する。
	_, f, _, _ := runtime.Caller(0)
	bs, err := os.ReadFile(filepath.Join(path.Dir(f), "..", "..", "..", "infra", "mysql", "schema.sql"))
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	for _, q := range strings.Split(string(bs), ";") {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}

		if _, err := client.db.Exec(q); err != nil {
			log.Fatalf("failed to execute query: %s", err)
		}
	}
}

// NewFixtureLoader はtestdataディレクトリ配下のyamlファイルを読み込むテストデータローダーを返す
func NewFixtureLoader(client *Client) *testfixtures.Loader {
	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Join(path.Dir(f), "..", "..", "testdata")
	fixtures, err := testfixtures.New(
		testfixtures.Database(client.db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(dir),
	)
	if err != nil {
		log.Fatalf("failed to create fixture loader: %s", err)
	}
	return fixtures
}
