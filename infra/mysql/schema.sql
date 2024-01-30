CREATE TABLE users (
    id         CHAR(26)     NOT NULL COMMENT 'ユーザID',
    name       VARCHAR(15)  NOT NULL COMMENT 'ユーザ名',
    email      VARCHAR(254) NOT NULL COMMENT 'メールアドレス',
    password   CHAR(60)     NOT NULL COMMENT 'パスワード',
    created_at DATETIME     NOT NULL COMMENT '作成日',
    updated_at DATETIME     NOT NULL COMMENT '更新日',
    PRIMARY KEY (id),
    UNIQUE INDEX (name),
    UNIQUE INDEX (email)
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'ユーザ';

CREATE TABLE projects (
    id          CHAR(26)    NOT NULL COMMENT 'プロジェクトID',
    user_id     CHAR(26)    NOT NULL COMMENT '所有するユーザのID',
    name        VARCHAR(20) NOT NULL COMMENT 'プロジェクト名',
    color       CHAR(7)     NOT NULL COMMENT 'カラー',
    is_archived BOOLEAN     NOT NULL COMMENT 'アーカイブされたか',
    created_at  DATETIME    NOT NULL COMMENT '作成日',
    updated_at  DATETIME    NOT NULL COMMENT '更新日',
    PRIMARY KEY (id),
    CONSTRAINT projects_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'プロジェクト';

CREATE TABLE tasks (
    id           CHAR(26)            NOT NULL COMMENT 'タスクID',
    user_id      CHAR(26)            NOT NULL COMMENT '所有するユーザのID',
    project_id   CHAR(26)            NOT NULL COMMENT '紐づくプロジェクトのID',
    title        VARCHAR(80)         NOT NULL COMMENT 'タイトル',
    content      VARCHAR(300)        NOT NULL COMMENT 'メモ',
    priority     TINYINT(2) UNSIGNED NOT NULL CHECK (priority BETWEEN 0 AND 3) COMMENT '優先度（0~3の数字で指定し、3が最も優先度が高い）',
    due_on       DATE COMMENT '期日',
    completed_at DATETIME COMMENT '完了日',
    created_at   DATETIME            NOT NULL COMMENT '作成日',
    updated_at   DATETIME            NOT NULL COMMENT '更新日',
    PRIMARY KEY (id),
    CONSTRAINT tasks_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT tasks_project_id_fk FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'タスク';

CREATE TABLE steps (
    id           CHAR(26)    NOT NULL COMMENT 'ステップID',
    user_id      CHAR(26)    NOT NULL COMMENT '所有するユーザのID',
    task_id      CHAR(26)    NOT NULL COMMENT '紐づくタスクのID',
    title        VARCHAR(80) NOT NULL COMMENT 'タイトル',
    completed_at DATETIME COMMENT '完了日',
    created_at   DATETIME    NOT NULL COMMENT '作成日',
    updated_at   DATETIME    NOT NULL COMMENT '更新日',
    PRIMARY KEY (id),
    CONSTRAINT steps_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT steps_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'ステップ';

CREATE TABLE tags (
    id         CHAR(26)    NOT NULL COMMENT 'タグID',
    user_id    CHAR(26)    NOT NULL COMMENT '所有するユーザのID',
    name       VARCHAR(20) NOT NULL COMMENT 'タグ名',
    created_at DATETIME    NOT NULL COMMENT '作成日',
    updated_at DATETIME    NOT NULL COMMENT '更新日',
    PRIMARY KEY (id),
    CONSTRAINT tags_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'タグ';

CREATE TABLE tasks_tags (
    task_id CHAR(26) NOT NULL COMMENT '紐づくタスクID',
    tag_id  CHAR(26) NOT NULL COMMENT '紐づくタグID',
    PRIMARY KEY (task_id, tag_id),
    CONSTRAINT tasks_tags_task_id_fk FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT tasks_tags_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'タスクとタグの紐づき';
