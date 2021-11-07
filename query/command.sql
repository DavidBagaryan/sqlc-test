-- name: CreateIntTableRow :exec
INSERT INTO int_status (status)
VALUES ($1)
RETURNING *;

-- name: CreateJsonbTableRow :exec
INSERT INTO jsonb_status (status)
VALUES ($1)
RETURNING *;

-- name: CreateRelationStatusRow :exec
INSERT INTO relation_status (status)
VALUES ($1)
RETURNING *;

-- name: CreateStatusDict :exec
INSERT INTO status_dictionary (name)
VALUES ($1)
RETURNING *;

-- name: PrepareTables :exec
TRUNCATE int_status, jsonb_status, relation_status CASCADE;
