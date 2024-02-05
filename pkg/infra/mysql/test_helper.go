package mysql

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/minguu42/simoom/pkg/infra/mysql/sqlc"
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
	dir := filepath.Join(path.Dir(f), "..", "..", "..", "testdata")
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

// InitAllData は全てのテーブルのデータを削除し、テストデータを投入し直す
// エラー時はパニックを投げる
func InitAllData(client *Client) {
	q := sqlc.New(client.db)
	ctx := context.Background()
	if err := q.DeleteAllUsers(ctx); err != nil {
		panic(err)
	}

	if err := q.ImportUser(ctx); err != nil {
		panic(err)
	}
	if err := q.ImportProject(ctx); err != nil {
		panic(err)
	}
	if err := q.ImportTag(ctx); err != nil {
		panic(err)
	}
	if err := q.ImportTask(ctx); err != nil {
		panic(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		panic(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		panic(err)
	}
}

// ResetProject はprojectテーブルのデータを削除し、テストデータを投入し直す
func ResetProject(t testing.TB, client *Client) {
	q := sqlc.New(client.db)
	ctx := context.Background()
	if err := q.DeleteAllProjects(ctx); err != nil {
		t.Fatal(err)
	}

	if err := q.ImportProject(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportTask(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		t.Fatal(err)
	}
}

// ResetStep はstepテーブルのデータを削除し、テストデータを投入し直す
func ResetStep(t testing.TB, client *Client) {
	q := sqlc.New(client.db)
	ctx := context.Background()
	if err := q.DeleteAllSteps(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		t.Fatal(err)
	}
}

// ResetTag はtagテーブルのデータを削除し、テストデータを投入し直す
// エラー時はパニックを投げる
func ResetTag(t testing.TB, client *Client) {
	q := sqlc.New(client.db)
	ctx := context.Background()
	if err := q.DeleteAllTags(ctx); err != nil {
		t.Fatal(err)
	}

	if err := q.ImportTag(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		t.Fatal(err)
	}
}

// ResetTask はtaskテーブルのデータを削除し、テストデータを投入し直す
// エラー時はパニックを投げる
func ResetTask(t testing.TB, client *Client) {
	q := sqlc.New(client.db)
	ctx := context.Background()
	if err := q.DeleteAllTasks(ctx); err != nil {
		t.Fatal(err)
	}

	if err := q.ImportTask(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportStep(ctx); err != nil {
		t.Fatal(err)
	}
	if err := q.ImportTaskTag(ctx); err != nil {
		t.Fatal(err)
	}
}
