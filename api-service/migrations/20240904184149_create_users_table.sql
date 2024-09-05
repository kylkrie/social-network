-- +goose Up
CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(15) NOT NULL,
    pfp_url TEXT,
    protected BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_users_username ON users(username);

CREATE TRIGGER update_user_modtime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_user_modtime ON users;
DROP INDEX IF EXISTS idx_users_username;
DROP TABLE IF EXISTS users;
