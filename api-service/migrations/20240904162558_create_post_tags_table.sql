-- +goose Up
CREATE TABLE post_tags (
    id BIGINT PRIMARY KEY,
    post_id BIGINT NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    start_index INTEGER,
    end_index INTEGER,
    tag TEXT
    --FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_post_tags_post_id ON post_tags(post_id);

-- +goose Down
DROP INDEX IF EXISTS idx_post_tags_post_id;
DROP TABLE IF EXISTS post_tags;
