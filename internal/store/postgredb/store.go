package postgredb

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ipr-savelichev/internal/config"
	"ipr-savelichev/internal/store"
	"ipr-savelichev/internal/store/postgredb/tables/employe"
	"ipr-savelichev/internal/store/postgredb/tables/task"
	"ipr-savelichev/internal/store/postgredb/tables/tool"
	"ipr-savelichev/internal/store/postgredb/tables/user"
	tableemploye "ipr-savelichev/internal/store/tables/employe"
	tabletask "ipr-savelichev/internal/store/tables/task"
	tabletool "ipr-savelichev/internal/store/tables/tool"
	tableuser "ipr-savelichev/internal/store/tables/user"
)

type Postgres struct {
	db                *gorm.DB
	logger            *zap.Logger
	taskRepository    tabletask.Task
	toolRepository    tabletool.Tool
	employeRepository tableemploye.Employe
	userRepository    tableuser.User
}

var (
	NoConnection = "Отсутсвует подключение к БД."
)

func NewGorm(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DataBase), &gorm.Config{})
	if err != nil {
		logger.Error("Ошибка подключения к БД.")
		return nil, err
	}

	logger.Info("Успешное подключение к БД.")

	return db, err
}

func NewPostgre(db *gorm.DB, logger *zap.Logger) (store.Store, error) {
	if db == nil {
		logger.Error(NoConnection)
		return nil, fmt.Errorf("%w", NoConnection)
	}

	return &Postgres{
		db:                db,
		logger:            logger,
		taskRepository:    task.NewTaskRepository(db, logger),
		toolRepository:    tool.NewToolRepository(db, logger),
		employeRepository: employe.NewEmployeRepository(db, logger),
		userRepository:    user.NewUserRepository(db, logger),
	}, nil
}

func (p *Postgres) Task() tabletask.Task {
	return p.taskRepository
}

func (p *Postgres) Tool() tabletool.Tool {
	return p.toolRepository
}

func (p *Postgres) Employe() tableemploye.Employe {
	return p.employeRepository
}

func (p *Postgres) User() tableuser.User {
	return p.userRepository
}
