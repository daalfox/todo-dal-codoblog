-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo (
    id serial PRIMARY KEY,
    title varchar NOT NULL,
    done boolean NOT NULL DEFAULT false
)
-- +goose StatementEnd
