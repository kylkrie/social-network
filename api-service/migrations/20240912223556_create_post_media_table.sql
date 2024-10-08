-- +goose Up
CREATE TABLE post_media (
    media_key BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    type VARCHAR(20) NOT NULL,
    url TEXT NOT NULL,
    width INT NOT NULL,
    height INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    --FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_post_media_post_id ON post_media(post_id);
CREATE INDEX idx_post_media_user_id ON post_media(user_id);

-- +goose Down
DROP INDEX IF EXISTS idx_post_media_post_id;
DROP INDEX IF EXISTS idx_post_media_user_id;
DROP TABLE IF EXISTS post_media;
