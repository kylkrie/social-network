-- +goose Up
CREATE TABLE post_likes (
    post_id BIGINT,
    user_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (post_id, user_id)
);

CREATE INDEX idx_post_likes_user_id ON post_likes(user_id);

-- +goose Down
DROP INDEX IF EXISTS idx_post_likes_user_id;
DROP TABLE IF EXISTS post_likes;
