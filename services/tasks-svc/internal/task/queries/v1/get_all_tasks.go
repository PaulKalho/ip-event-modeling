package v1

import (
	"context"
	pb "gen/services/tasks_svc/v1"
	"github.com/google/uuid"
	"hwdb"
	"tasks-svc/internal/task/models"
	"tasks-svc/repos/task_repo"
)

type GetAllTasksWithSubtasksQueryHandler func(ctx context.Context) ([]*models.Task, error)

func NewGetAllTasksWithSubtasksQueryHandler() GetAllTasksWithSubtasksQueryHandler {
	return func(ctx context.Context) ([]*models.Task, error) {
		taskRepo := task_repo.New(hwdb.GetDB())

		rows, err := taskRepo.GetAllTasks(ctx)
		if err := hwdb.Error(ctx, err); err != nil {
			return nil, err
		}

		tasks := make([]*models.Task, 0)
		tasksMap := make(map[uuid.UUID]int)

		for _, row := range rows {
			var task *models.Task

			if ix, exists := tasksMap[row.Task.ID]; exists {
				task = tasks[ix]
			} else {
				task = &models.Task{
					ID:          row.Task.ID,
					Name:        row.Task.Name,
					Description: row.Task.Description,
					Status:      pb.TaskStatus(row.Task.Status),
					Public:      row.Task.Public,
					DueAt:       nil, // may be set below
					CreatedBy:   row.Task.CreatedBy,
					CreatedAt:   row.Task.CreatedAt.Time,
					Subtasks:    make(map[uuid.UUID]models.Subtask),
				}
				if row.Task.DueAt.Valid {
					task.DueAt = &row.Task.DueAt.Time
				}

				tasks = append(tasks, task)
				tasksMap[row.Task.ID] = len(tasks) - 1
			}

			if row.Task.DueAt.Valid {
				task.Subtasks[row.SubtaskID.UUID] = models.Subtask{
					ID:   row.SubtaskID.UUID,
					Name: *row.SubtasksName,
					Done: *row.SubtasksDone,
				}
			}
		}

		return tasks, nil
	}
}
