-- +goose Up
CREATE TABLE post_references (
    id BIGINT PRIMARY KEY,
    source_post_id BIGINT NOT NULL,
    referenced_post_id BIGINT NOT NULL,
    reference_type VARCHAR(50) NOT NULL
    --FOREIGN KEY (source_post_id) REFERENCES posts(id) ON DELETE CASCADE,
    --FOREIGN KEY (referenced_post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE INDEX idx_post_references_source_post_id ON post_references(source_post_id);
CREATE INDEX idx_post_references_referenced_post_id ON post_references(referenced_post_id);

-- +goose Down
DROP INDEX IF EXISTS idx_post_references_referenced_post_id;
DROP INDEX IF EXISTS idx_post_references_source_post_id;
DROP TABLE IF EXISTS post_references;
