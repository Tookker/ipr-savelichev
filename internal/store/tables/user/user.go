package tableuser

import (
	"context"
	"ipr-savelichev/internal/models/user"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=User
type User interface {
	Login(context.Context, *user.User) error
	Register(context.Context, *user.User) error
}
