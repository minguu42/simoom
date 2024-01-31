-- name: ImportUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES ('user_01', 'ユーザ1', 'user01@example.com', 'some-password','2020-01-01 00:00:01', '2020-01-01 00:00:01');

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: ImportProject :exec
INSERT INTO projects (id, user_id, name, color, is_archived, created_at, updated_at)
VALUES ('project_01', 'user_01', 'プロジェクト1', '#1a2b3c', FALSE, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('project_02', 'user_01', 'プロジェクト2', '#ffffff', FALSE, '2020-01-01 00:00:02', '2020-01-01 00:00:02');

-- name: DeleteAllProjects :exec
DELETE FROM projects;

-- name: ImportStep :exec
INSERT INTO steps (id, user_id, task_id, title, completed_at, created_at, updated_at)
VALUES ('step_01', 'user_01', 'task_01', 'ステップ1', NULL, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('step_02', 'user_01', 'task_01', 'ステップ2', NULL, '2020-01-01 00:00:02', '2020-01-01 00:00:02');

-- name: DeleteAllSteps :exec
DELETE FROM steps;

-- name: ImportTag :exec
INSERT INTO tags (id, user_id, name, created_at, updated_at)
VALUES ('tag_01', 'user_01', 'タグ1', '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('tag_02', 'user_01', 'タグ2', '2020-01-01 00:00:02', '2020-01-01 00:00:02');

-- name: DeleteAllTags :exec
DELETE FROM tags;

-- name: ImportTask :exec
INSERT INTO tasks (id, user_id, project_id, title, content, priority, due_on, completed_at, created_at, updated_at)
VALUES ('task_01', 'user_01', 'project_01', 'タスク1', 'コンテンツ1', 3, '2020-01-02', NULL, '2020-01-01 00:00:01', '2020-01-01 00:00:01'),
       ('task_2', 'user_01', 'project_01', 'タスク2', '', 0, NULL, NULL, '2020-01-01 00:00:02', '2020-01-01 00:00:02');

-- name: DeleteAllTasks :exec
DELETE FROM tasks;

-- name: ImportTaskTag :exec
INSERT INTO tasks_tags (task_id, tag_id)
VALUES ('task_01', 'tag_01'),
       ('task_01', 'tag_02');

-- name: DeleteAllTaskTags :exec
DELETE FROM tasks_tags;
