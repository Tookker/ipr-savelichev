package user

import (
	"net/http"

	"go.uber.org/zap"

	"ipr-savelichev/internal/jwt"
	"ipr-savelichev/internal/models/coder"
	"ipr-savelichev/internal/models/user"
	"ipr-savelichev/internal/router/returnvalue"
	"ipr-savelichev/internal/store"
)

// TODO добавить тесты
type User interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type UserControl struct {
	store  store.Store
	jwt    jwt.JWT
	logger *zap.Logger
}

func NewUserController(store store.Store, jwt jwt.JWT, logger *zap.Logger) User {
	return &UserControl{
		store:  store,
		jwt:    jwt,
		logger: logger,
	}
}

func (u *UserControl) Login(w http.ResponseWriter, r *http.Request) {
	user, err := coder.Decode[user.User](r)
	if err != nil {
		u.logger.Error("Ошибка декодирования структуры user.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = u.store.User().Login(r.Context(), &user)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	jwt, err := u.jwt.GenerateJWT(user.Login)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, jwt)
}

func (u *UserControl) Register(w http.ResponseWriter, r *http.Request) {
	user, err := coder.Decode[user.User](r)
	if err != nil {
		u.logger.Error("Ошибка декодирования структуры user.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = u.store.User().Register(r.Context(), &user)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}
