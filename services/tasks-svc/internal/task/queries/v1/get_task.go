package v1

import (
	"context"
	"hwes"
	"tasks-svc/internal/task/aggregate"
	"tasks-svc/internal/task/models"

	"github.com/google/uuid"
)

type GetTaskQueryHandler func(ctx context.Context, taskID uuid.UUID) (*models.Task, error)

func NewGetTaskQueryHandler(as hwes.AggregateStore) GetTaskQueryHandler {
	return func(ctx context.Context, taskID uuid.UUID) (*models.Task, error) {
		task, err := aggregate.LoadTaskAggregate(ctx, as, taskID)
		if err != nil {
			return nil, err
		}

		return task.Task, nil
	}
}
