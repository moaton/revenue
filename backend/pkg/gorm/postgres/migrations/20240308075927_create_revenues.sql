-- +goose Up
-- +goose StatementBegin

CREATE TABLE revenue_groups(
  id uuid NOT NULL PRIMARY KEY,
  user_id uuid NOT NULL,
  name varchar NOT NULL,
  type varchar,
  amount bigint NOT NULL,
  description text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE revenues(
  id uuid NOT NULL PRIMARY KEY,
  user_id uuid NOT NULL,
  name varchar NOT NULL,
  type varchar NOT NULL,
  amount bigint NOT NULL,
  description text,
  group_id uuid,
  datetime timestamp,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (group_id) REFERENCES revenue_groups(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE revenues;
DROP TABLE revenue_groups;
-- +goose StatementEnd
