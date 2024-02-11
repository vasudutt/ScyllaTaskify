CREATE TABLE IF NOT EXISTS todos (
	id          UUID PRIMARY KEY,
	user_id     TEXT,
	title       TEXT,
	description TEXT,
	status      INT,
	created_at  TIMESTAMP,
	updated_at  TIMESTAMP
);

CREATE INDEX ON todos (user_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS todos_byuser_id AS
    SELECT id, user_id, title, description, status, created_at, updated_at
    FROM todos
    WHERE user_id IS NOT NULL AND id IS NOT NULL
    PRIMARY KEY (user_id, id);

CREATE MATERIALIZED VIEW IF NOT EXISTS todos_by_status AS
    SELECT id, user_id, title, description, status, created_at, updated_at
    FROM todos
    WHERE status IS NOT NULL AND id IS NOT NULL
    PRIMARY KEY (status, id);
