package employe

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/employe"
	"ipr-savelichev/internal/models/task"
	tableemploye "ipr-savelichev/internal/store/tables/employe"
)

const (
	tableEmployess = "employess"
	tableTask      = "tasks"
)

type employeRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// Конструктор
func NewEmployeRepository(db *gorm.DB, logger *zap.Logger) tableemploye.Employe {
	return &employeRepository{
		db:     db,
		logger: logger,
	}
}

// Получение списка работников из БД
func (r *employeRepository) GetAllEmployes(ctx context.Context) ([]employe.Employe, error) {
	var employes []employe.Employe
	res := r.db.WithContext(ctx).Table(tableEmployess).Find(&employes)
	if res.Error != nil {
		r.logger.Error("Ошибка получения записей из БД. Описание: " + res.Error.Error())
		return nil, res.Error
	}

	r.logger.Info("Успешное добавление записи в БД.")

	return employes, nil
}

func (r *employeRepository) GetTaskEmploye(ctx context.Context, id uint) (task.Task, error) {
	var task task.Task
	res := r.db.WithContext(ctx).Table(tableTask).Where("employeid = ?", id).Find(&task)
	if res.Error != nil {
		r.logger.Error("Ошибка получения записей из БД. Описание: " + res.Error.Error())
		return task, res.Error
	}

	return task, nil
}

// Удаление работника из БД
func (r *employeRepository) RemoveEmploye(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Table(tableEmployess).Delete(&employe.Employe{Id: id})
	if res.Error != nil {
		r.logger.Error("Ошибка удаления записи из БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное удаление записи в БД.")

	return nil
}

// Изменить данные работника в БД
func (r *employeRepository) EditEmploye(ctx context.Context, emp *employe.Employe) error {
	res := r.db.WithContext(ctx).Table(tableEmployess).Where("id = ?", emp.Id).Updates(&employe.Employe{
		Name: emp.Name,
		Age:  emp.Age,
		Sex:  emp.Sex,
	})
	if res.Error != nil {
		r.logger.Error("Ошибка изменения записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное изменение записи в БД.")

	return nil
}

// Добавть нового работника в БД
func (r *employeRepository) AddEmploye(ctx context.Context, employe *employe.Employe) error {
	res := r.db.WithContext(ctx).Table(tableEmployess).Create(employe)
	if res.Error != nil {
		r.logger.Error("Ошибка добавления записи в БД. Описание: " + res.Error.Error())
		return res.Error
	}

	r.logger.Info("Успешное добавление записи в БД.")

	return nil
}
