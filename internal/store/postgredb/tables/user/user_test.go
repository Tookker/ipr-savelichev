package user_test

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ipr-savelichev/internal/models/user"
	userRepository "ipr-savelichev/internal/store/postgredb/tables/user"
	"ipr-savelichev/internal/store/tables/user/mocks"
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

// Login(context.Context, *user.User) error
func TestLogin(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	userRep := userRepository.NewUserRepository(db, logger)

	type args struct {
		ctx  context.Context
		user *user.User
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
			name: "Login user",
			args: args{
				ctx: context.Background(),
				user: &user.User{
					Login:    "user",
					Password: "user",
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
			repo := &mocks.User{}
			repo.On("Login", tt.args.ctx).Return(tt.want)

			err := userRep.Login(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// Register(context.Context, *user.User) error
func TestRegister(t *testing.T) {
	logger, err := getZapLogger()
	if err != nil {
		return
	}

	db, err := connectDB(dsn)
	if err != nil {
		return
	}

	userRep := userRepository.NewUserRepository(db, logger)

	type args struct {
		ctx  context.Context
		user *user.User
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
			name: "Register user",
			args: args{
				ctx: context.Background(),
				user: &user.User{
					Login:    "user",
					Password: "user",
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
			repo := &mocks.User{}
			repo.On("Register", tt.args.ctx).Return(tt.want)

			err := userRep.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
