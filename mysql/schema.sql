CREATE TABLE user (
  id         CHAR(26)    NOT NULL,
  name       VARCHAR(15) NOT NULL UNIQUE,
  api_key    CHAR(64)    NOT NULL UNIQUE,
  created_at DATETIME    NOT NULL,
  updated_at DATETIME    NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE project (
  id          CHAR(26)    NOT NULL,
  user_id     CHAR(26)    NOT NULL,
  name        VARCHAR(20) NOT NULL,
  color       CHAR(7)     NOT NULL,
  is_archived BOOLEAN     NOT NULL,
  created_at  DATETIME    NOT NULL,
  updated_at  DATETIME    NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT project_user_id_fk FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE task (
  id           CHAR(26)     NOT NULL,
  user_id      CHAR(26)     NOT NULL,
  project_id   CHAR(26)     NOT NULL,
  title        VARCHAR(80)  NOT NULL,
  content      VARCHAR(300) NOT NULL,
  priority     TINYINT(2) UNSIGNED NOT NULL CHECK (priority BETWEEN 0 AND 3),
  due_on       DATE,
  completed_at DATETIME,
  created_at   DATETIME     NOT NULL,
  updated_at   DATETIME     NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT task_user_id_fk FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT task_project_id_fk FOREIGN KEY (project_id) REFERENCES project(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE step (
  id           CHAR(26)    NOT NULL,
  user_id      CHAR(26)    NOT NULL,
  task_id      CHAR(26)    NOT NULL,
  title        VARCHAR(80) NOT NULL,
  completed_at DATETIME,
  created_at   DATETIME    NOT NULL,
  updated_at   DATETIME    NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT step_user_id_fk FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT step_task_id_fk FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE tag (
  id         CHAR(26)    NOT NULL,
  user_id    CHAR(26)    NOT NULL,
  name       VARCHAR(20) NOT NULL,
  created_at DATETIME    NOT NULL,
  updated_at DATETIME    NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT tag_user_id_fk FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE task_tag (
  task_id CHAR(26) NOT NULL,
  tag_id  CHAR(26) NOT NULL,
  PRIMARY KEY (task_id, tag_id),
  CONSTRAINT task_tag_task_id_fk FOREIGN KEY (task_id) REFERENCES task (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT task_tag_tag_id_fk FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE ON UPDATE CASCADE
);
