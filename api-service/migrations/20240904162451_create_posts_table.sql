-- +goose Up
CREATE TABLE posts (
    id BIGINT PRIMARY KEY,
    content TEXT,
    author_id BIGINT,
    conversation_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_posts_author_id ON posts(author_id);
CREATE INDEX idx_posts_conversation_id ON posts(conversation_id);
CREATE INDEX idx_posts_created_at_brin ON posts USING BRIN (created_at) WITH (pages_per_range = 128);
CREATE INDEX idx_posts_content_gin ON posts USING gin(to_tsvector('english', content));

CREATE TRIGGER update_post_modtime
    BEFORE UPDATE ON posts
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_post_modtime ON posts;
DROP INDEX IF EXISTS idx_posts_content_gin;
DROP INDEX IF EXISTS idx_posts_created_at_brin;
DROP INDEX IF EXISTS idx_posts_conversation_id;
DROP INDEX IF EXISTS idx_posts_author_id;
DROP TABLE IF EXISTS posts;
