-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS int_status
(
    id     uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    status int         NOT NULL
);

-- CREATE INDEX IF NOT EXISTS int_status ON int_status (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS int_status;
-- +goose StatementEnd
