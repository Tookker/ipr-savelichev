package user

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/user"
	tableuser "ipr-savelichev/internal/store/tables/user"
)

const (
	userTableName = "users"
)

var (
	ErrUserExist = errors.New("Пользователь существует!")
)

type userRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) tableuser.User {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (u *userRepository) Login(ctx context.Context, user *user.User) error {
	res := u.db.WithContext(ctx).Table(userTableName).Take(&user)
	if res.Error != nil {
		u.logger.Error("Ошибка получения данных пользователя из БД. Описание: " + res.Error.Error())
		return res.Error
	}

	return nil
}

func (u *userRepository) Register(ctx context.Context, usr *user.User) error {
	res := u.db.WithContext(ctx).Table(userTableName).Transaction(func(tx *gorm.DB) error {
		var tempUsr user.User
		tx.Where("login = ?", usr.Login).Find(&tempUsr)
		if tempUsr == *usr {
			u.logger.Warn("Пользователь существует!")
			return ErrUserExist
		}

		tx.Create(&usr)
		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	if res != nil {
		u.logger.Error("Ошибка записи пользователя в БД. Описание: " + res.Error())
		return res
	}

	return nil
}
