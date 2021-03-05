-- +goose Up
CREATE INDEX images_user_idx ON images (user_id);

-- +goose Down
DROP INDEX IF EXISTS images_user_idx;