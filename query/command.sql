-- name: CreateIntTableRow :exec
INSERT INTO int_status (status)
VALUES ($1)
RETURNING *;

-- name: CreateJsonbTableRow :exec
INSERT INTO jsonb_status (status)
VALUES ($1)
RETURNING *;

-- name: PrepareTables :exec
TRUNCATE int_status, jsonb_status CASCADE;
