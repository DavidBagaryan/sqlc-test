-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS status_dictionary
(
    id   SERIAL PRIMARY KEY,
    name varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS relation_status
(
    id     uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    status int         NOT NULL REFERENCES status_dictionary (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS relation_status, status_dictionary;
-- +goose StatementEnd
