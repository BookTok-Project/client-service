CREATE TABLE IF NOT EXISTS user_favorite_cards (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    card_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (user_id, card_id)
);
