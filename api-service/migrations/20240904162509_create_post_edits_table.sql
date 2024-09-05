-- +goose Up
CREATE TABLE post_edits (
    id BIGINT PRIMARY KEY,
    post_id BIGINT NOT NULL,
    content TEXT,
    edited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    --FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_post_edits_post_id ON post_edits(post_id);

-- +goose Down
DROP INDEX IF EXISTS idx_post_edits_post_id;
DROP TABLE IF EXISTS post_edits;
