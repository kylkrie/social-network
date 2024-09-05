-- +goose Up
CREATE TABLE user_public_metrics (
    user_id BIGINT PRIMARY KEY,
    followers_count INT DEFAULT 0,
    following_count INT DEFAULT 0,
    post_count INT DEFAULT 0,
    listed_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_user_metrics_modtime
    BEFORE UPDATE ON user_public_metrics
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_user_metrics_modtime ON user_public_metrics;
DROP TABLE IF EXISTS user_public_metrics;
