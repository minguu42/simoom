CREATE TABLE users (
    id         CHAR(26)     NOT NULL COMMENT 'ユーザID',
    name       VARCHAR(15)  NOT NULL UNIQUE,
    email      VARCHAR(254) NOT NULL UNIQUE,
    password   CHAR(60)     NOT NULL UNIQUE,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE projects (
    id          CHAR(26)    NOT NULL,
    user_id     CHAR(26)    NOT NULL,
    name        VARCHAR(20) NOT NULL,
    color       CHAR(7)     NOT NULL,
    is_archived BOOLEAN     NOT NULL,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT projects_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE tasks (
    id           CHAR(26)            NOT NULL,
    user_id      CHAR(26)            NOT NULL,
    project_id   CHAR(26)            NOT NULL,
    title        VARCHAR(80)         NOT NULL,
    content      VARCHAR(300)        NOT NULL,
    priority     TINYINT(2) UNSIGNED NOT NULL CHECK (priority BETWEEN 0 AND 3),
    due_on       DATE,
    completed_at DATETIME,
    created_at   DATETIME            NOT NULL,
    updated_at   DATETIME            NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT tasks_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT tasks_project_id_fk FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE steps (
    id           CHAR(26)    NOT NULL,
    user_id      CHAR(26)    NOT NULL,
    task_id      CHAR(26)    NOT NULL,
    title        VARCHAR(80) NOT NULL,
    completed_at DATETIME,
    created_at   DATETIME    NOT NULL,
    updated_at   DATETIME    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT steps_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT steps_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE tags (
    id         CHAR(26)    NOT NULL,
    user_id    CHAR(26)    NOT NULL,
    name       VARCHAR(20) NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT tags_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE tasks_tags (
    task_id CHAR(26) NOT NULL,
    tag_id  CHAR(26) NOT NULL,
    PRIMARY KEY (task_id, tag_id),
    CONSTRAINT tasks_tags_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT tasks_tags_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
