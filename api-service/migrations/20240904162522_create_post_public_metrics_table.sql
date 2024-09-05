-- +goose Up
CREATE TABLE post_public_metrics (
    post_id BIGINT PRIMARY KEY,
    repost_count INT DEFAULT 0,
    reply_count INT DEFAULT 0,
    like_count INT DEFAULT 0,
    quote_count INT DEFAULT 0,
    view_count INT DEFAULT 0
    --FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS post_public_metrics;
