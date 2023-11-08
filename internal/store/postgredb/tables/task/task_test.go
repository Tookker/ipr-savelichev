package task_test

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/task"
	taskRepository "ipr-savelichev/internal/store/postgredb/tables/task"
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

// AddTask(context.Context, *task.Task) error
func TestAddTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	taskRep := taskRepository.NewTaskRepository(db, logger)

	type args struct {
		ctx  context.Context
		task task.Task
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
				task: task.Task{
					Id:          4,
					Descryption: "Строить забор",
					EmployeID:   1,
					ToolID:      2,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("AddTask", tt.args.ctx, tt.args.task).Return(tt.want)

			err := taskRep.AddTask(tt.args.ctx, &tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// GetAllTask(context.Context) ([]task.EmployeTask, error)
func TestGetAllTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	taskRep := taskRepository.NewTaskRepository(db, logger)

	type args struct {
		ctx context.Context
	}

	type want struct {
		tasks []task.EmployeTask
		err   error
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
				tasks: []task.EmployeTask{
					{
						Id:              1,
						DescryptionTask: "Копать огород",
						Name:            "Яичко Лиза Сергеевна",
						Age:             "2000-11-21T00:00:00Z",
						Sex:             "Ж",
						DescryptionTool: "Лопата",
					},
					{
						Id:              4,
						DescryptionTask: "Строить забор",
						Name:            "Петров Петр Петрович",
						Age:             "1976-02-03T00:00:00Z",
						Sex:             "М",
						DescryptionTool: "Молоток",
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
			repo.On("GetAllTask", tt.args.ctx).Return(tt.want)

			tasks, err := taskRep.GetAllTask(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}

			if len(tasks) != len(tt.want.tasks) {
				t.Errorf("GetAllTasks error! Whant %v, get %v ", tt.want.tasks, tasks)
				return
			}

			for i := 0; i < len(tasks); i++ {
				if tasks[i] != tt.want.tasks[i] {
					t.Errorf("GetAllTasks error! Whant %v, get %v ", tt.want.tasks, tasks)
					return
				}
			}
		})
	}
}

// GetTask(context.Context, uint) (task.EmployeTask, error)
func TestGetTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	taskRep := taskRepository.NewTaskRepository(db, logger)

	type args struct {
		ctx context.Context
		id  uint
	}

	type want struct {
		task *task.EmployeTask
		err  error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Get task",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: want{
				task: &task.EmployeTask{
					Id:              1,
					DescryptionTask: "Копать огород",
					Name:            "Яичко Лиза Сергеевна",
					Age:             "2000-11-21T00:00:00Z",
					Sex:             "Ж",
					DescryptionTool: "Лопата",
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

			task, err := taskRep.GetTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}

			if task != *tt.want.task {
				t.Errorf("GetTask error whant = %v, get %v", *tt.want.task, task)
				return
			}
		})
	}
}

// EditTask(context.Context, *task.Task) error
func TestEditTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	taskRep := taskRepository.NewTaskRepository(db, logger)

	type args struct {
		ctx  context.Context
		task *task.Task
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
				task: &task.Task{
					Id:          1,
					Descryption: "Вскапывать огород",
					EmployeID:   3,
					ToolID:      1,
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
			repo.On("EditTask", tt.args.ctx, tt.args.task).Return(tt.want)

			err := taskRep.EditTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditTask error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// RemoveTask(context.Context, uint) error
func TestRemoveTask(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	taskRep := taskRepository.NewTaskRepository(db, logger)

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
			name: "Remove task",
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
			name: "Remove dont exiset field task",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want: want{
				err: nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("RemoveTask", tt.args.ctx, tt.args.id).Return(tt.want)

			err := taskRep.RemoveTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
