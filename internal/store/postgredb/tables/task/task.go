package task

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/task"
	tabletask "ipr-savelichev/internal/store/tables/task"
)

const (
	tableName = "tasks"
)

var (
	ErrRecrodNotFound = errors.New("record not found")
)

type taskRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTaskRepository(db *gorm.DB, logger *zap.Logger) tabletask.Task {
	return &taskRepository{
		db:     db,
		logger: logger,
	}
}

func (r *taskRepository) GetAllTask(ctx context.Context) ([]task.EmployeTask, error) {
	var tasks []task.EmployeTask

	res := r.db.WithContext(ctx).
		Table(tableName).
		Scopes(getTaskEmployeScope).
		Scan(&tasks)
	if res.Error != nil {
		r.logger.Error("Ошибка получения записей в БД. Описание: " + res.Error.Error())
		return nil, res.Error
	}

	r.logger.Info("Успешное получение записей в БД.")

	return tasks, nil
}

func (r *taskRepository) AddTask(ctx context.Context, tsk *task.Task) error {
	res := r.db.WithContext(ctx).Table(tableName).Create(tsk)

	if res.Error != nil {
		r.logger.Error("Ошибка добавление записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное добавление записи в БД.")

	return nil
}

func (r *taskRepository) RemoveTask(ctx context.Context, id uint) error {
	task := task.Task{Id: id}
	res := r.db.WithContext(ctx).Table(tableName).Take(&task).Delete(&task)
	if res.Error != nil {
		r.logger.Error("Ошибка удаления записи из БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное удаление записи в БД.")

	return nil
}

func (r *taskRepository) EditTask(ctx context.Context, tsk *task.Task) error {
	res := r.db.WithContext(ctx).Table(tableName).Where("id = ?", tsk.Id).Updates(&task.Task{
		Descryption: tsk.Descryption,
		EmployeID:   tsk.EmployeID,
		ToolID:      tsk.ToolID,
	})

	if res.Error != nil {
		r.logger.Error("Ошибка изменения записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное изменение записи в БД.")

	return nil
}

func (r *taskRepository) GetTask(ctx context.Context, id uint) (task.EmployeTask, error) {
	var task task.EmployeTask
	res := r.db.WithContext(ctx).
		Table(tableName).
		Scopes(getTaskEmployeScope).
		Where("tasks.id = ?", id).
		Scan(&task)

	if res.Error != nil {
		r.logger.Error("Ошибка получения записи в БД. Описание: " + res.Error.Error())
		return task, res.Error
	}

	r.logger.Info("Успешное получение записи в БД.")

	return task, nil
}

func getTaskEmployeScope(db *gorm.DB) *gorm.DB {
	return db.Select("tasks.id, tasks.descryption as descryptiontask, employess.name, employess.age, employess.sex, tool.descryption as descryptiontool").
		Joins("join employess on tasks.employeid=employess.id").
		Joins("join tool on tasks.toolid=tool.id")
}
