-- +goose Up
CREATE TABLE images (
    "email" VARCHAR(50),
    "gcp_id" VARCHAR(50),
    "id" SERIAL NOT NULL,
    "price" FLOAT,
    "title" VARCHAR(150),
    "url" VARCHAR(100),
    "user_id" VARCHAR(50), 
    PRIMARY KEY(id)
);
-- +goose Down
DROP TABLE images