-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS jsonb_status
(
    id     uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    status jsonb       NOT NULL
);

-- CREATE INDEX IF NOT EXISTS jsonb_status ON jsonb_status (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS jsonb_status;
-- +goose StatementEnd
