package employe

import (
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"ipr-savelichev/internal/models/coder"
	"ipr-savelichev/internal/models/employe"
	"ipr-savelichev/internal/router/returnvalue"
	"ipr-savelichev/internal/store"
)

const (
	debugMsg = "Получен запрос: "
	errID    = "Неверный идентификатор!"
)

type Employe interface {
	GetAllEmployes(w http.ResponseWriter, r *http.Request)
	GetTaskEmploye(w http.ResponseWriter, r *http.Request)
	RemoveEmploye(w http.ResponseWriter, r *http.Request)
	EditEmploye(w http.ResponseWriter, r *http.Request)
	AddEmploye(w http.ResponseWriter, r *http.Request)
}

type EmployeControl struct {
	store  store.Store
	logger *zap.Logger
}

// Конструктор
func NewEmployeControl(store store.Store, logger *zap.Logger) Employe {
	return &EmployeControl{
		store:  store,
		logger: logger,
	}
}

// Handler обработки запроса на получение списка работников
func (e *EmployeControl) GetAllEmployes(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug(debugMsg + "получение списка работников из таблицы Employes.")

	res, err := e.store.Employe().GetAllEmployes(r.Context())
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, res)
}

// TODO реализовать хендлер после реализация обработки запросов для Tasks
func (e *EmployeControl) GetTaskEmploye(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug(debugMsg + "получение задачи рабоника из таблицы Task.")
	const idPath = 2

	res := strings.Split(r.URL.String(), "/")

	getID, err := strconv.Atoi(res[idPath])
	if err != nil || getID <= 0 {
		e.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	task, err := e.store.Employe().GetTaskEmploye(r.Context(), uint(getID))
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, task)
}

// Handler обработки запроса на удаление строки из списка работников
func (e *EmployeControl) RemoveEmploye(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug(debugMsg + "удаление работника из таблицы Employes.")

	res := strings.Split(r.URL.String(), "/")
	deleteID, err := strconv.Atoi(res[len(res)-1])
	if err != nil || deleteID <= 0 {
		e.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	err = e.store.Employe().RemoveEmploye(r.Context(), uint(deleteID))
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}

// Handler обработки запроса на изменение строки из списка работников
func (e *EmployeControl) EditEmploye(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug(debugMsg + "изменение данных работника из таблицы Employes.")

	employe, err := coder.Decode[employe.Employe](r)
	if err != nil {
		e.logger.Error("Ошибка декодирования структуры employe.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = e.store.Employe().EditEmploye(r.Context(), &employe)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}

// Handler обработки запроса на добавление нового работника
func (e *EmployeControl) AddEmploye(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug(debugMsg + "добавление данных нового работника в таблицу Employes.")

	employe, err := coder.Decode[employe.Employe](r)
	if err != nil {
		e.logger.Error("Ошибка декодирования структуры employe.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = e.store.Employe().AddEmploye(r.Context(), &employe)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}
