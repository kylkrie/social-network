-- +goose Up
CREATE TABLE user_profiles (
    user_id BIGINT PRIMARY KEY,
    banner_url TEXT,
    bio TEXT,
    website TEXT,
    location VARCHAR(100),
    birthday DATE,
    pinned_post_id BIGINT,
    followers INT DEFAULT 0,
    following INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_user_profile_modtime
    BEFORE UPDATE ON user_profiles
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_user_profile_modtime ON user_profiles;
DROP TABLE IF EXISTS user_profiles;
