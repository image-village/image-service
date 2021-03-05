-- +goose Up
ALTER TABLE images
ADD CONSTRAINT user_id_unique UNIQUE (user_id);

-- +goose Down
ALTER TABLE images
DROP CONSTRAINT user_id_unique;