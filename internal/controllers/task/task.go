package task

import (
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"ipr-savelichev/internal/models/coder"
	"ipr-savelichev/internal/models/task"
	"ipr-savelichev/internal/router/returnvalue"
	"ipr-savelichev/internal/store"
)

const (
	debugMsg = "Получен запрос: "
	errID    = "Неверный идентификатор!"
)

type Task interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	GetAllTask(w http.ResponseWriter, r *http.Request)
	AddTask(w http.ResponseWriter, r *http.Request)
	EditTask(w http.ResponseWriter, r *http.Request)
	RemoveTask(w http.ResponseWriter, r *http.Request)
}

type TaskControl struct {
	store  store.Store
	logger *zap.Logger
}

func NewTaskControl(store store.Store, logger *zap.Logger) Task {
	return &TaskControl{
		store:  store,
		logger: logger,
	}
}

// Hanler получение задачи из БД Task
func (t *TaskControl) GetTask(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug(debugMsg + "изменение элемента в таблице Tasks.")
	res := strings.Split(r.URL.String(), "/")
	getID, err := strconv.Atoi(res[len(res)-1])
	if err != nil || getID <= 0 {
		t.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	task, err := t.store.Task().GetTask(r.Context(), uint(getID))
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, task)
}

// Hanler для добавления задачи в БД Task
func (t *TaskControl) AddTask(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug(debugMsg + "добавление элемента в таблицу Task.")

	task, err := coder.Decode[task.Task](r)
	if err != nil {
		t.logger.Error("Ошибка декодирования структуры task.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = t.store.Task().AddTask(r.Context(), &task)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}

// Hanler для изменения задачи в БД Task
func (t *TaskControl) EditTask(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug(debugMsg + "изменение элемента в таблице Tasks.")
	res := strings.Split(r.URL.String(), "/")
	editID, err := strconv.Atoi(res[len(res)-1])
	if err != nil || editID <= 0 {
		t.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	task, err := coder.Decode[task.Task](r)
	if err != nil {
		t.logger.Error("Ошибка декодирования структуры task.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = t.store.Task().EditTask(r.Context(), &task)
	if err != nil {
		errDisc := err.Error()
		t.logger.Error(errDisc)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errDisc)
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)

}

// Hanler удаление задачи из БД Task
func (t *TaskControl) RemoveTask(w http.ResponseWriter, r *http.Request) {
	t.logger.Debug(debugMsg + "удаление элемента из таблицы Inventory.")
	res := strings.Split(r.URL.String(), "/")
	deleteID, err := strconv.Atoi(res[len(res)-1])
	if err != nil || deleteID <= 0 {
		t.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	err = t.store.Task().RemoveTask(r.Context(), uint(deleteID))
	if err != nil {
		errDesc := err.Error()
		t.logger.Error(errDesc)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errDesc)
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)

}

// Hanler получение всех задач из БД Task
func (t *TaskControl) GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.store.Task().GetAllTask(r.Context())
	if err != nil {
		errDesc := err.Error()
		t.logger.Error(errDesc)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errDesc)
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, tasks)
}
