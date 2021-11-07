// Code generated by sqlc. DO NOT EDIT.
// source: command.sql

package repository

import (
	"context"
	"encoding/json"
)

const createIntTableRow = `-- name: CreateIntTableRow :exec
INSERT INTO int_status (status)
VALUES ($1)
RETURNING id, status
`

func (q *Queries) CreateIntTableRow(ctx context.Context, status int32) error {
	_, err := q.db.ExecContext(ctx, createIntTableRow, status)
	return err
}

const createJsonbTableRow = `-- name: CreateJsonbTableRow :exec
INSERT INTO jsonb_status (status)
VALUES ($1)
RETURNING id, status
`

func (q *Queries) CreateJsonbTableRow(ctx context.Context, status json.RawMessage) error {
	_, err := q.db.ExecContext(ctx, createJsonbTableRow, status)
	return err
}

const createRelationStatusRow = `-- name: CreateRelationStatusRow :exec
INSERT INTO relation_status (status)
VALUES ($1)
RETURNING id, status
`

func (q *Queries) CreateRelationStatusRow(ctx context.Context, status int32) error {
	_, err := q.db.ExecContext(ctx, createRelationStatusRow, status)
	return err
}

const createStatusDict = `-- name: CreateStatusDict :exec
INSERT INTO status_dictionary (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateStatusDict(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createStatusDict, name)
	return err
}

const prepareTables = `-- name: PrepareTables :exec
TRUNCATE int_status, jsonb_status, relation_status CASCADE
`

func (q *Queries) PrepareTables(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, prepareTables)
	return err
}
