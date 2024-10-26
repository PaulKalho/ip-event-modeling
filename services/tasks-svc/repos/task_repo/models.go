// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package task_repo

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SchemaMigration struct {
	Version int64
	Dirty   bool
}

type Subtask struct {
	ID     uuid.UUID
	TaskID uuid.UUID
	Name   string
	Done   bool
}

type Task struct {
	ID          uuid.UUID
	Name        string
	Description string
	Status      int32
	Public      bool
	CreatedBy   uuid.UUID
	DueAt       pgtype.Timestamp
	CreatedAt   pgtype.Timestamp
}
