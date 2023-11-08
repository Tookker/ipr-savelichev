package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"ipr-savelichev/internal/controllers/employe"
	"ipr-savelichev/internal/controllers/task"
	"ipr-savelichev/internal/controllers/tool"
	"ipr-savelichev/internal/controllers/user"
	"ipr-savelichev/internal/jwt"
	"ipr-savelichev/internal/store"
)

type ChiRouter struct {
	router *chi.Mux
	logger *zap.Logger
	store  store.Store
	jwt    jwt.JWT
}

func NewRouter(logger *zap.Logger, store store.Store, jwt jwt.JWT) *ChiRouter {
	return &ChiRouter{
		router: chi.NewRouter(),
		logger: logger,
		store:  store,
		jwt:    jwt,
	}
}

func (r *ChiRouter) SetHandlers(store store.Store) {
	r.setLoggerMidleware()
	r.setTaskHandlers(store)
	r.setEmployeHandlers(store)
	r.setToolHandlers(store)
	r.setUserHandlers(store)
}

func (r *ChiRouter) setTaskHandlers(store store.Store) {
	var handlers = task.NewTaskControl(r.store, r.logger)

	r.router.Route("/task", func(rout chi.Router) {
		rout.Use(r.jwt.MiddlewareJWT)
		rout.Get("/", handlers.GetAllTask)
		rout.Get("/{task_id}", handlers.GetTask)
		rout.Delete("/{task_id}", handlers.RemoveTask)
		rout.Patch("/{task_id}", handlers.EditTask)
		rout.Post("/", handlers.AddTask)
	})
}

func (r *ChiRouter) setEmployeHandlers(store store.Store) {
	var handler = employe.NewEmployeControl(r.store, r.logger)

	r.router.Route("/employe", func(rout chi.Router) {
		rout.Use(r.jwt.MiddlewareJWT)
		rout.Get("/", handler.GetAllEmployes)
		rout.Post("/", handler.AddEmploye)
		rout.Delete("/{employe_id}", handler.RemoveEmploye)
		rout.Patch("/{employe_id}", handler.EditEmploye)
		rout.Get("/{employe_id}/task", handler.GetTaskEmploye)
	})
}

func (r *ChiRouter) setToolHandlers(store store.Store) {
	var handler = tool.NewToolControl(r.store, r.logger)

	r.router.Route("/tools", func(rout chi.Router) {
		rout.Use(r.jwt.MiddlewareJWT)
		rout.Get("/", handler.GetAllTools)
		rout.Delete("/{tools_id}", handler.RemoveTool)
		rout.Patch("/", handler.EditTool)
		rout.Post("/", handler.AddTool)
	})
}

func (r *ChiRouter) setUserHandlers(store store.Store) {
	var handler = user.NewUserController(r.store, r.jwt, r.logger)
	r.router.Get("/login", handler.Login)
	r.router.Post("/register", handler.Register)
}

func (r *ChiRouter) setLoggerMidleware() {
	r.router.Use(middleware.Logger)
	r.router.Use(middleware.RequestID)
}

func (r *ChiRouter) GetRouter() *chi.Mux {
	return r.router
}
