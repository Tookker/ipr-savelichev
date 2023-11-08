package tabletask

import (
	"context"
	"ipr-savelichev/internal/models/task"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=Task
type Task interface {
	GetAllTask(context.Context) ([]task.EmployeTask, error)
	GetTask(context.Context, uint) (task.EmployeTask, error)
	AddTask(context.Context, *task.Task) error
	RemoveTask(context.Context, uint) error
	EditTask(context.Context, *task.Task) error
}
