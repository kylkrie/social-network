-- +goose Up
CREATE TABLE post_public_metrics (
    post_id BIGINT PRIMARY KEY,
    reposts INT DEFAULT 0,
    replies INT DEFAULT 0,
    likes INT DEFAULT 0,
    views INT DEFAULT 0
    --FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS post_public_metrics;
