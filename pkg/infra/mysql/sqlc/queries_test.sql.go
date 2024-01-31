// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries_test.sql

package sqlc

import (
	"context"
)

const deleteAllProjects = `-- name: DeleteAllProjects :exec
DELETE FROM projects
`

func (q *Queries) DeleteAllProjects(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllProjects)
	return err
}

const deleteAllSteps = `-- name: DeleteAllSteps :exec
DELETE FROM steps
`

func (q *Queries) DeleteAllSteps(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllSteps)
	return err
}

const deleteAllTags = `-- name: DeleteAllTags :exec
DELETE FROM tags
`

func (q *Queries) DeleteAllTags(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllTags)
	return err
}

const deleteAllTaskTags = `-- name: DeleteAllTaskTags :exec
DELETE FROM tasks_tags
`

func (q *Queries) DeleteAllTaskTags(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllTaskTags)
	return err
}

const deleteAllTasks = `-- name: DeleteAllTasks :exec
DELETE FROM tasks
`

func (q *Queries) DeleteAllTasks(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllTasks)
	return err
}

const deleteAllUsers = `-- name: DeleteAllUsers :exec
DELETE FROM users
`

func (q *Queries) DeleteAllUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllUsers)
	return err
}

const importProject = `-- name: ImportProject :exec
INSERT INTO projects (id, user_id, name, color, is_archived, created_at, updated_at)
VALUES ('project_01', 'user_01', 'プロジェクト1', '#1a2b3c', FALSE, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('project_02', 'user_01', 'プロジェクト2', '#ffffff', FALSE, '2020-01-01 00:00:02', '2020-01-01 00:00:02')
`

func (q *Queries) ImportProject(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importProject)
	return err
}

const importStep = `-- name: ImportStep :exec
INSERT INTO steps (id, user_id, task_id, name, completed_at, created_at, updated_at)
VALUES ('step_01', 'user_01', 'task_01', 'ステップ1', NULL, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('step_02', 'user_01', 'task_01', 'ステップ2', NULL, '2020-01-01 00:00:02', '2020-01-01 00:00:02')
`

func (q *Queries) ImportStep(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importStep)
	return err
}

const importTag = `-- name: ImportTag :exec
INSERT INTO tags (id, user_id, name, created_at, updated_at)
VALUES ('tag_01', 'user_01', 'タグ1', '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('tag_02', 'user_01', 'タグ2', '2020-01-01 00:00:02', '2020-01-01 00:00:02')
`

func (q *Queries) ImportTag(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importTag)
	return err
}

const importTask = `-- name: ImportTask :exec
INSERT INTO tasks (id, user_id, project_id, name, content, priority, due_on, completed_at, created_at, updated_at)
VALUES ('task_01', 'user_01', 'project_01', 'タスク1', 'コンテンツ1', 3, '2020-01-02', NULL, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('task_2', 'user_01', 'project_01', 'タスク2', '', 0, NULL, NULL, '2020-01-01 00:00:02', '2020-01-01 00:00:02')
`

func (q *Queries) ImportTask(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importTask)
	return err
}

const importTaskTag = `-- name: ImportTaskTag :exec
INSERT INTO tasks_tags (task_id, tag_id)
VALUES ('task_01', 'tag_01'),
       ('task_01', 'tag_02')
`

func (q *Queries) ImportTaskTag(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importTaskTag)
	return err
}

const importUser = `-- name: ImportUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES ('user_01', 'ユーザ1', 'user01@example.com', 'some-password','2020-01-01 00:00:01', '2020-01-01 00:00:01')
`

func (q *Queries) ImportUser(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, importUser)
	return err
}
