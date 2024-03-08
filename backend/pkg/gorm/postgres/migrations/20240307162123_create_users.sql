-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
  id uuid NOT NULL PRIMARY KEY,
  username varchar NOT NULL UNIQUE,
  password varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
