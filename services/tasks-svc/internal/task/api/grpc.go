package api

import (
	"context"
	pb "gen/services/tasks_svc/v1"
	"hwes"
	"hwutil"

	"github.com/google/uuid"
	"tasks-svc/internal/task/handlers"
)

type TaskGrpcService struct {
	pb.UnimplementedTaskServiceServer
	as       hwes.AggregateStore
	handlers *handlers.Handlers
}

func NewTaskGrpcService(aggregateStore hwes.AggregateStore, handlers *handlers.Handlers) *TaskGrpcService {
	return &TaskGrpcService{as: aggregateStore, handlers: handlers}
}

func (s *TaskGrpcService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	taskID := uuid.New()
	subtasks := hwutil.OrEmptySlice(req.GetSubtasks())

	err := s.handlers.Commands.V1.CreateTask(
		ctx,
		taskID,
		req.GetName(),
		req.Description,
		req.Public,
		req.InitialStatus,
		req.DueAt,
		subtasks,
	)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Id: taskID.String(),
	}, nil
}

func (s *TaskGrpcService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	taskID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	err = s.handlers.Commands.V1.UpdateTask(
		ctx, taskID, req.Name, req.Description, req.Status, req.Public, req.DueAt)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{}, nil
}

func (s *TaskGrpcService) CreateSubtask(
	ctx context.Context,
	req *pb.CreateSubtaskRequest,
) (*pb.CreateSubtaskResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, err
	}

	subtaskID := uuid.New()

	err = s.handlers.Commands.V1.CreateSubtask(
		ctx, taskID, subtaskID, req.GetSubtask().GetName(), req.GetSubtask().GetDone())
	if err != nil {
		return nil, err
	}

	return &pb.CreateSubtaskResponse{
		SubtaskId: subtaskID.String(),
	}, nil
}

func (s *TaskGrpcService) UpdateSubtask(
	ctx context.Context,
	req *pb.UpdateSubtaskRequest,
) (*pb.UpdateSubtaskResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, err
	}

	subtaskID, err := uuid.Parse(req.GetSubtaskId())
	if err != nil {
		return nil, err
	}

	err = s.handlers.Commands.V1.UpdateSubtask(
		ctx, taskID, subtaskID, req.Subtask.Name, req.Subtask.Done)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateSubtaskResponse{}, nil
}

func (s *TaskGrpcService) DeleteSubtask(
	ctx context.Context,
	req *pb.DeleteSubtaskRequest,
) (*pb.DeleteSubtaskResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, err
	}

	subtaskID, err := uuid.Parse(req.GetSubtaskId())
	if err != nil {
		return nil, err
	}

	err = s.handlers.Commands.V1.DeleteSubtask(ctx, taskID, subtaskID)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSubtaskResponse{}, nil
}

func (s *TaskGrpcService) RemoveTaskDueDate(
	ctx context.Context,
	req *pb.RemoveTaskDueDateRequest,
) (*pb.RemoveTaskDueDateResponse, error) {
	taskID, err := uuid.Parse(req.GetTaskId())
	if err != nil {
		return nil, err
	}

	err = s.handlers.Commands.V1.RemoveTaskDueAt(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveTaskDueDateResponse{}, nil
}

func (s *TaskGrpcService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	taskID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	if err := s.handlers.Commands.V1.DeleteTask(ctx, taskID); err != nil {
		return nil, err
	}

	return &pb.DeleteTaskResponse{}, nil
}
