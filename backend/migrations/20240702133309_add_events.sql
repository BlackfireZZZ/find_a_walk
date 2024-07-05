-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    author_id UUID NOT NUll,
    start_longitude FLOAT NOT NULL,
    start_latitude FLOAT NOT NULL,
    end_longitude FLOAT NOT NULL,
    end_latitude FLOAT NOT NULL,
    date DATE NOT NULL,
    capacity INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS members (
    id TEXT PRIMARY KEY,
    event_id UUID NOT NULL,
    user_id UUID NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS event_tags (
    event_id UUID NOT NULL,
    tag_id TEXT NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (event_id, tag_id)
);


-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS events CASCADE;
DROP TABLE IF EXISTS members CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS event_tags CASCADE;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
