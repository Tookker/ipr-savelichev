package tableemploye

import (
	"context"

	"ipr-savelichev/internal/models/employe"
	"ipr-savelichev/internal/models/task"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=Employe
type Employe interface {
	GetAllEmployes(context.Context) ([]employe.Employe, error)
	GetTaskEmploye(context.Context, uint) (task.Task, error)
	RemoveEmploye(context.Context, uint) error
	EditEmploye(context.Context, *employe.Employe) error
	AddEmploye(context.Context, *employe.Employe) error
}
