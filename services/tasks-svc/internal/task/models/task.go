package models

import (
	pb "gen/services/tasks_svc/v1"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	CreatedBy uuid.UUID
	CreatedAt time.Time

	Name        string
	Description string
	Status      pb.TaskStatus
	Public      bool
	DueAt       *time.Time
	Subtasks    map[uuid.UUID]Subtask
}

type Subtask struct {
	ID        uuid.UUID
	CreatedBy uuid.UUID
	CreatedAt time.Time

	Name string
	Done bool
}
