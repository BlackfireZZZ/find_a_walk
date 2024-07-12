-- +goose Up
-- +goose StatementBegin
INSERT INTO tags (id, name)
SELECT 'sex', 'sex' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'sex');
INSERT INTO tags (id, name)
SELECT 'programming', 'programming' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'programming');
INSERT INTO tags (id, name)
SELECT 'music', 'music' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'music');
INSERT INTO tags (id, name)
SELECT 'football', 'football' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'football');
INSERT INTO tags (id, name)
SELECT 'history', 'history' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'history');
INSERT INTO tags (id, name)
SELECT 'gaming', 'gaming' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'gaming');
INSERT INTO tags (id, name)
SELECT 'travel', 'travel' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'travel');
INSERT INTO tags (id, name)
SELECT 'fitness', 'fitness' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'fitness');
INSERT INTO tags (id, name)
SELECT 'movies', 'movies' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'movies');
INSERT INTO tags (id, name)
SELECT 'photography', 'photography' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'photography');
INSERT INTO tags (id, name)
SELECT 'reading', 'reading' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'reading');
INSERT INTO tags (id, name)
SELECT 'art', 'art' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'art');
INSERT INTO tags (id, name)
SELECT 'cooking', 'cooking' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'cooking');
INSERT INTO tags (id, name)
SELECT 'fashion', 'fashion' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'fashion');
INSERT INTO tags (id, name)
SELECT 'technology', 'technology' WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 'technology');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tags WHERE id IN ('sex', 'programming', 'music', 'football', 'history', 'gaming', 'travel', 'fitness', 'movies', 'photography', 'reading', 'art', 'cooking', 'fashion', 'technology');

-- +goose StatementEnd

