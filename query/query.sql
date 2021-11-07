-- name: GetCountIntTableByStatus :one
SELECT count(*)
FROM int_status
WHERE status & $1 != 0;

-- name: GetCountJsonbTableByStatus :one
SELECT count(*)
FROM jsonb_status
WHERE status ? $1::text;

-- name: GetCountRelationTableByStatus :one
SELECT count(*)
FROM relation_status;


