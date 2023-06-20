package task_template

import (
	"context"
	pb "gen/proto/services/task_svc/v1"
	"github.com/google/uuid"
	zlog "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hwgorm"
	"hwutil"
)

type Base struct {
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

type TaskTemplate struct {
	Base
	ID       uuid.UUID             `gorm:"column:id"`
	Public   bool                  `gorm:"column:is_public"`
	UserID   *uuid.UUID            `gorm:"column:user_id;default:NULL"`
	WardID   *uuid.UUID            `gorm:"column:ward_id;default:NULL"`
	SubTasks []TaskTemplateSubtask `gorm:"foreignKey:TaskTemplateID"`
}

type TaskTemplateSubtask struct {
	ID             uuid.UUID `gorm:"column:id"`
	TaskTemplateID uuid.UUID `gorm:"column:task_template_id"`
	Name           string    `gorm:"column:name"`
}

type ServiceServer struct {
	pb.UnimplementedTaskTemplateServiceServer
}

func NewServiceServer() *ServiceServer {
	return &ServiceServer{}
}

func (ServiceServer) CreateTaskTemplate(ctx context.Context, req *pb.CreateTaskTemplateRequest) (*pb.CreateTaskTemplateResponse, error) {
	log := zlog.Ctx(ctx)
	db := hwgorm.GetDB(ctx)

	var userID uuid.UUID
	if req.UserId != nil {
		var err error
		userID, err = uuid.Parse(*req.UserId)
		if err != nil {
			return nil, err
		}
	}

	var wardID uuid.UUID
	if req.WardId != nil {
		var err error
		wardID, err = uuid.Parse(*req.WardId)
		if err != nil {
			return nil, err
		}
	}

	/*
		if req.WardId == nil && req.UserId == nil {
			return nil, status.Error(codes.InvalidArgument, "either wardID or userID must be set")
		}*/

	// authenticate the userID or wardID

	taskTemplate := TaskTemplate{
		Base:   Base{Name: req.Name, Description: req.Description},
		Public: req.Public,
		UserID: &userID,
		WardID: &wardID,
	}

	if err := db.Create(&taskTemplate).Error; err != nil {
		log.Warn().Err(err).Msg("database error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	var subtasks []TaskTemplateSubtask
	if req.Subtasks != nil {
		subtasks = hwutil.Map(req.Subtasks, func(subtask *pb.CreateTaskTemplateRequest_SubTask) TaskTemplateSubtask {
			return TaskTemplateSubtask{Name: subtask.Name, TaskTemplateID: taskTemplate.ID}
		})
	}

	taskTemplate.SubTasks = append(taskTemplate.SubTasks, subtasks...)
	if err := db.Save(&taskTemplate).Error; err != nil {
		log.Warn().Err(err).Msg("database error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info().
		Str("taskTemplateId", taskTemplate.ID.String()).
		Msg("taskTemplate created")

	return &pb.CreateTaskTemplateResponse{
		Id: taskTemplate.ID.String(),
	}, nil
}
