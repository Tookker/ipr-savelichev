package tabletool

import (
	"context"
	"ipr-savelichev/internal/models/tool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=Tool
type Tool interface {
	GetAllTools(context.Context) ([]tool.Tool, error)
	AddTool(context.Context, *tool.Tool) error
	RemoveTool(context.Context, uint) error
	EditTool(context.Context, *tool.Tool) error
}
