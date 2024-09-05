-- +goose Up
CREATE TABLE user_relationships (
    user_id BIGINT,
    related_user_id BIGINT,
    relationship_type SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, related_user_id, relationship_type)
);

CREATE INDEX idx_user_relationships_user_type ON user_relationships(user_id, relationship_type);
CREATE INDEX idx_user_relationships_related_type ON user_relationships(related_user_id, relationship_type);

-- +goose Down
DROP INDEX IF EXISTS idx_user_relationships_related_type;
DROP INDEX IF EXISTS idx_user_relationships_user_type;
DROP TABLE IF EXISTS user_relationships;
