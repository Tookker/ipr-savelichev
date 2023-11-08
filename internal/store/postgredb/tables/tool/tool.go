package tool

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/tool"
	tabletool "ipr-savelichev/internal/store/tables/tool"
)

type toolRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewToolRepository(db *gorm.DB, logger *zap.Logger) tabletool.Tool {
	return &toolRepository{
		db:     db,
		logger: logger,
	}
}

func (r *toolRepository) GetAllTools(ctx context.Context) ([]tool.Tool, error) {
	var tools []tool.Tool
	res := r.db.WithContext(ctx).Table("tool").Find(&tools)
	if res.Error != nil {
		r.logger.Error("Ошибка получения записей из БД. Описание: " + res.Error.Error())
		return nil, res.Error
	}

	r.logger.Info("Успешное получение записей из БД.")

	return tools, nil
}

func (r *toolRepository) AddTool(ctx context.Context, tool *tool.Tool) error {
	res := r.db.WithContext(ctx).Table("tool").Create(tool)
	if res.Error != nil {
		r.logger.Error("Ошибка добавления записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное добавление записи в БД.")

	return nil
}

func (r *toolRepository) RemoveTool(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Table("tool").Take(&tool.Tool{Id: id}).Delete(&tool.Tool{Id: id})
	if res.Error != nil {
		r.logger.Error("Ошибка удаления записи из БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное удаление записи из БД.")

	return nil
}

func (r *toolRepository) EditTool(ctx context.Context, tool *tool.Tool) error {
	res := r.db.WithContext(ctx).Table("tool").Where("id = ?", tool.Id).Update("descryption", tool.Descryption)
	if res.Error != nil {
		r.logger.Error("Ошибка изменения записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное изменение записи в БД.")

	return nil
}
