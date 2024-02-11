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
