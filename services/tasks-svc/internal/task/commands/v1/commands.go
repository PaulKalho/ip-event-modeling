package v1

type TaskCommands struct {
	CreateTask      CreateTaskCommandHandler
	UpdateTask      UpdateTaskCommandHandler
	CreateSubtask   CreateSubtaskCommandHandler
	UpdateSubtask   UpdateSubtaskCommandHandler
	DeleteSubtask   DeleteSubtaskCommandHandler
	RemoveTaskDueAt RemoveTaskDueAtCommandHandler
	DeleteTask      DeleteTaskCommandHandler
}
