DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id         CHAR(26) PRIMARY KEY,
  name       VARCHAR(15) NOT NULL UNIQUE,
  api_key    CHAR(64)    NOT NULL UNIQUE,
  created_at DATETIME    NOT NULL,
  updated_at DATETIME    NOT NULL
);

DROP TABLE IF EXISTS projects;
CREATE TABLE projects (
  id          CHAR(26) PRIMARY KEY,
  user_id     CHAR(26)    NOT NULL,
  name        VARCHAR(20) NOT NULL,
  color       CHAR(7)     NOT NULL,
  is_archived BOOLEAN     NOT NULL,
  created_at  DATETIME    NOT NULL,
  updated_at  DATETIME    NOT NULL
);

DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks (
  id           CHAR(26) PRIMARY KEY,
  project_id   CHAR(26)            NOT NULL,
  title        VARCHAR(80)         NOT NULL,
  content      VARCHAR(300)        NOT NULL,
  priority     TINYINT(2) UNSIGNED NOT NULL,
  due_on       DATE,
  completed_at DATETIME,
  created_at   DATETIME            NOT NULL,
  updated_at   DATETIME            NOT NULL
);

DROP TABLE IF EXISTS steps;
CREATE TABLE steps (
  id           CHAR(26) PRIMARY KEY,
  task_id      CHAR(26)    NOT NULL,
  title        VARCHAR(80) NOT NULL,
  completed_at DATETIME,
  created_at   DATETIME    NOT NULL,
  updated_at   DATETIME    NOT NULL
);

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
  id   CHAR(26) PRIMARY KEY,
  name VARCHAR(20) NOT NULL
);

DROP TABLE IF EXISTS tasks_tags;
CREATE TABLE tasks_tags (
  task_id CHAR(26) NOT NULL,
  tag_id  CHAR(26) NOT NULL,
  PRIMARY KEY (task_id, tag_id)
);
