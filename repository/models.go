// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"encoding/json"

	"github.com/google/uuid"
)

type IntStatus struct {
	ID     uuid.UUID
	Status int32
}

type JsonbStatus struct {
	ID     uuid.UUID
	Status json.RawMessage
}
