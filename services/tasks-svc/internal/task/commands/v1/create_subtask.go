package v1

import (
	"context"
	"hwes"

	"github.com/google/uuid"

	"tasks-svc/internal/task/aggregate"
)

type CreateSubtaskCommandHandler func(
	ctx context.Context,
	taskID,
	subtaskID uuid.UUID,
	name string,
	done bool,
) error

func NewCreateSubtaskCommandHandler(as hwes.AggregateStore) CreateSubtaskCommandHandler {
	return func(
		ctx context.Context,
		taskID,
		subtaskID uuid.UUID,
		name string,
		done bool,
	) error {
		a, err := aggregate.LoadTaskAggregate(ctx, as, taskID)
		if err != nil {
			return err
		}

		if err := a.CreateSubtask(ctx, subtaskID, name, done); err != nil {
			return err
		}

		return as.Save(ctx, a)
	}
}
