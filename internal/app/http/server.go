package server

import (
	"context"
	"net/http"
	tasksHandler "todos-api/internal/transport/http/tasks"
	tasksUsecase "todos-api/internal/usecase/tasks"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func New(
	addr string,
	taskUC tasksUsecase.UseCase,
) *Server {
	srv := &http.Server{
		Addr: addr,
	}

	th := tasksHandler.New(taskUC)

	r := configureRouter(th)

	srv.Handler = r

	return &Server{
		srv: srv,
	}
}

func configureRouter(th *tasksHandler.Handler) *gin.Engine {
	r := gin.Default()

	tasksGroup := r.Group("/tasks")
	tasksHandler.RegisterRoutes(tasksGroup, th)

	return r
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
