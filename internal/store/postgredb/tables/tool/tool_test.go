package tool_test

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/tool"
	toolRepository "ipr-savelichev/internal/store/postgredb/tables/tool"
	"ipr-savelichev/internal/store/tables/tool/mocks"
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

// GetAllTools(ctx context.Context) ([]tool.Tool, error)
func TestGetAllTools(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	toolRep := toolRepository.NewToolRepository(db, logger)

	type args struct {
		ctx context.Context
	}

	type want struct {
		tools []tool.Tool
		err   error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "Get all tools from BD",
			args:    args{ctx: context.Background()},
			want:    want{tools: []tool.Tool{{Id: 1, Descryption: "Лопата"}, {Id: 2, Descryption: "Молоток"}, {Id: 3, Descryption: "Метла"}}, err: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Tool{}
			repo.On("GetAllTools", tt.args.ctx).Return(tt.want)

			tools, err := toolRep.GetAllTools(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTools error = %v, whantErr %v", err, tt.wantErr)
				return
			}

			if len(tools) != len(tt.want.tools) {
				t.Errorf("GetAllTools error! Whant %v, get %v ", tt.want.tools, tools)
				return
			}

			for i := 0; i < len(tools); i++ {
				if tools[i] != tt.want.tools[i] {
					t.Errorf("GetAllTools error! Whant %v, get %v ", tt.want.tools, tools)
					return
				}
			}
		})
	}
}

// AddTool(ctx context.Context, tool *tool.Tool) error
func TestAddTool(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	toolRep := toolRepository.NewToolRepository(db, logger)

	type args struct {
		ctx  context.Context
		tool *tool.Tool
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
			name:    "Add tool to BD",
			args:    args{ctx: context.Background(), tool: &tool.Tool{Descryption: "Вилы"}},
			want:    want{err: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Tool{}
			repo.On("AddTool", tt.args.ctx, tt.args.tool).Return(tt.want)

			err := toolRep.AddTool(tt.args.ctx, tt.args.tool)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTool error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// RemoveTool(ctx context.Context, id uint) error
func TestRemoveTool(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	toolRep := toolRepository.NewToolRepository(db, logger)

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
			name:    "Remove tool from BD",
			args:    args{ctx: context.Background(), id: 2},
			want:    want{err: nil},
			wantErr: false,
		},
		{
			name:    "Remove dont exist tool from BD",
			args:    args{ctx: context.Background(), id: 4},
			want:    want{err: nil},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Tool{}
			repo.On("RemoveTool", tt.args.ctx, tt.args.id).Return(tt.want)

			err := toolRep.RemoveTool(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveTool error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// EditTool(ctx context.Context, tool *tool.Tool) error
func TestEditTool(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	toolRep := toolRepository.NewToolRepository(db, logger)

	type args struct {
		ctx  context.Context
		tool *tool.Tool
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
			name:    "Edit tool from BD",
			args:    args{ctx: context.Background(), tool: &tool.Tool{Id: 1, Descryption: "Вилы"}},
			want:    want{err: nil},
			wantErr: false,
		},
		{
			name:    "Edit dont exist tool from BD",
			args:    args{ctx: context.Background(), tool: &tool.Tool{Id: 4, Descryption: "Вилы"}},
			want:    want{err: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Tool{}
			repo.On("EditTool", tt.args.ctx, tt.args.tool).Return(tt.want)

			err := toolRep.EditTool(tt.args.ctx, tt.args.tool)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditTool error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
