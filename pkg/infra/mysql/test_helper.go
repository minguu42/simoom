package mysql

import (
	"context"
	"testing"

	"github.com/minguu42/simoom/pkg/infra/mysql/sqlc"
)

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
