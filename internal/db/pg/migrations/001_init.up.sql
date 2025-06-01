CREATE TABLE IF NOT EXISTS subscriptions (
	id            BIGSERIAL PRIMARY KEY,
	subscriber_id BIGINT    NOT NULL,
	subscribee_id BIGINT    NOT NULL
);