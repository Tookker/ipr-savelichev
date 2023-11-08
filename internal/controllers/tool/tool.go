package tool

import (
	"net/http"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"ipr-savelichev/internal/models/coder"
	"ipr-savelichev/internal/models/tool"
	"ipr-savelichev/internal/router/returnvalue"
	"ipr-savelichev/internal/store"
)

const (
	debugMsg = "Получен запрос: "
)

type Tool interface {
	GetAllTools(w http.ResponseWriter, r *http.Request)
	AddTool(w http.ResponseWriter, r *http.Request)
	RemoveTool(w http.ResponseWriter, r *http.Request)
	EditTool(w http.ResponseWriter, r *http.Request)
}

type ToolControl struct {
	store  store.Store
	logger *zap.Logger
}

func NewToolControl(store store.Store, logger *zap.Logger) Tool {
	return &ToolControl{
		store:  store,
		logger: logger,
	}
}

func (t *ToolControl) GetAllTools(w http.ResponseWriter, r *http.Request) {
	tools, err := t.store.Tool().GetAllTools(r.Context())
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOKWithValue(http.StatusOK, &tools)
}

func (t *ToolControl) AddTool(w http.ResponseWriter, r *http.Request) {
	tool, err := coder.Decode[tool.Tool](r)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = t.store.Tool().AddTool(r.Context(), &tool)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}

func (t *ToolControl) RemoveTool(w http.ResponseWriter, r *http.Request) {
	const (
		errID = "Неверный идентификатор!"
	)

	t.logger.Debug(debugMsg + "удаление элемента из таблицы Inventory.")
	res := strings.Split(r.URL.String(), "/")
	deleteID, err := strconv.Atoi(res[len(res)-1])
	if err != nil || deleteID <= 0 {
		t.logger.Error(errID)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errID)
		return
	}

	err = t.store.Tool().RemoveTool(r.Context(), uint(deleteID))
	if err != nil {
		errDesc := err.Error()
		t.logger.Error(errDesc)
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, errDesc)
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}

func (t *ToolControl) EditTool(w http.ResponseWriter, r *http.Request) {
	tool, err := coder.Decode[tool.Tool](r)
	if err != nil {
		t.logger.Error("Ошибка декодирования структуры tool.")
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	err = t.store.Tool().EditTool(r.Context(), &tool)
	if err != nil {
		returnvalue.NewResponse(w).ResponseErr(http.StatusBadRequest, err.Error())
		return
	}

	returnvalue.NewResponse(w).ResponseOK(http.StatusOK)
}
