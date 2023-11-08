package employe_test

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/employe"
	"ipr-savelichev/internal/models/task"
	employeRepository "ipr-savelichev/internal/store/postgredb/tables/employe"
	"ipr-savelichev/internal/store/tables/task/mocks"
)

const (
	dsn = "host=127.0.0.1 user=admin password=root dbname=Tasks port=5432"
)

func getZapLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("Err create logger.")
		return nil, err
	}

	return logger, nil
}

func connectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Err open db.")
		return nil, err
	}

	return db, nil
}

// AddEmploye(context.Context, *employe.Employe) error
func TestAddEmploye(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	employeRep := employeRepository.NewEmployeRepository(db, logger)

	type args struct {
		ctx     context.Context
		employe *employe.Employe
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "Add correct task",
			args: args{
				ctx: context.Background(),
				employe: &employe.Employe{
					Name: "Алексеев Петр Андреевич",
					Age:  "1976-02-02",
					Sex:  "М",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("AddEmploye", tt.args.ctx, tt.args.employe).Return(tt.want)

			err := employeRep.AddEmploye(tt.args.ctx, tt.args.employe)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddEmploye error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// GetAllEmployes(context.Context) ([]employe.Employe, error)
func TestGetAllEmployes(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	employeRep := employeRepository.NewEmployeRepository(db, logger)

	type args struct {
		ctx context.Context
	}

	type want struct {
		employes []employe.Employe
		err      error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Get all task",
			args: args{
				ctx: context.Background(),
			},
			want: want{
				employes: []employe.Employe{
					{
						Id:   1,
						Name: "Петров Петр Петрович",
						Age:  "1976-02-03T00:00:00Z",
						Sex:  "М",
					},
					{
						Id:   2,
						Name: "Иванов Иван Иванович",
						Age:  "1997-04-15T00:00:00Z",
						Sex:  "М",
					},
					{
						Id:   3,
						Name: "Яичко Лиза Сергеевна",
						Age:  "2000-11-21T00:00:00Z",
						Sex:  "Ж",
					},
					{
						Id:   4,
						Name: "Алексеев Петр Андреевич",
						Age:  "1976-02-02T00:00:00Z",
						Sex:  "М",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("GetAllEmployes", tt.args.ctx).Return(tt.want)

			employes, err := employeRep.GetAllEmployes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}

			if len(employes) != len(tt.want.employes) {
				t.Errorf("GetAllTasks error! Whant %v, get %v ", tt.want.employes, employes)
				return
			}

			for i := 0; i < len(employes); i++ {
				if employes[i] != tt.want.employes[i] {
					t.Errorf("GetAllTasks error! Whant %v, get %v ", tt.want.employes, employes)
					return
				}
			}
		})
	}
}

// GetTaskEmploye(context.Context, uint) (task.Task, error)
func TestTaskEmploye(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	employeRep := employeRepository.NewEmployeRepository(db, logger)

	type args struct {
		ctx context.Context
		id  uint
	}

	type want struct {
		task task.Task
		err  error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "get task employe",
			args: args{
				ctx: context.Background(),
				id:  3,
			},
			want: want{
				task: task.Task{
					Id:          1,
					Descryption: "Вскапывать огород",
					EmployeID:   3,
					ToolID:      1,
				},
				err: nil,
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("GetTask", tt.args.ctx, tt.args.id).Return(tt.want)

			task, err := employeRep.GetTaskEmploye(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}

			if task != tt.want.task {
				t.Errorf("GetTask error whant = %v, get %v", tt.want.task, task)
				return
			}
		})
	}
}

// EditEmploye(context.Context, *employe.Employe) error
func TestEditTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	employeRep := employeRepository.NewEmployeRepository(db, logger)

	type args struct {
		ctx     context.Context
		employe *employe.Employe
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Edit task",
			args: args{
				ctx: context.Background(),
				employe: &employe.Employe{
					Id:   1,
					Name: "Петр Петр Петр",
					Age:  "1976-02-03T00:00:00Z",
					Sex:  "М",
				},
			},
			want: want{
				err: nil,
			},

			wantErr: false,
		},

		{
			name: "Edit dont exist task",
			args: args{
				ctx: context.Background(),
				employe: &employe.Employe{
					Id:   5,
					Name: "Петр Петр Петр",
					Age:  "1976-02-03T00:00:00Z",
					Sex:  "М",
				},
			},
			want: want{
				err: nil,
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("EditTask", tt.args.ctx, tt.args.employe).Return(tt.want)

			err := employeRep.EditEmploye(tt.args.ctx, tt.args.employe)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// RemoveEmploye(context.Context, uint) error
func TestRemoveEmploye(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	employeRep := employeRepository.NewEmployeRepository(db, logger)

	type args struct {
		ctx context.Context
		id  uint
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Remove employe",
			args: args{
				ctx: context.Background(),
				id:  4,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Remove dont exiset field employe",
			args: args{
				ctx: context.Background(),
				id:  5,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("RemoveTask", tt.args.ctx, tt.args.id).Return(tt.want)

			err := employeRep.RemoveEmploye(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
