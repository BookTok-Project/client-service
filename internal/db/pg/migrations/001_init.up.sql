CREATE TABLE IF NOT EXISTS subscriptions (
	id           BIGSERIAL PRIMARY KEY,
	reader_id    BIGINT    NOT NULL,
	writer_id    BIGINT    NOT NULL
);