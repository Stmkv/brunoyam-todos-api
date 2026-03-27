package server

import (
	"context"
	"net/http"
	authHandler "todos-api/internal/transport/http/auth"
	"todos-api/internal/transport/http/middleware"
	tasksHandler "todos-api/internal/transport/http/tasks"
	usersHandler "todos-api/internal/transport/http/users"
	authUseCase "todos-api/internal/usecase/auth"
	tasksUsecase "todos-api/internal/usecase/tasks"
	usersUsecase "todos-api/internal/usecase/users"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "todos-api/docs" //
)

type Server struct {
	srv *http.Server
}

func New(
	addr string,
	taskUC tasksUsecase.UseCase,
	userUC usersUsecase.UseCase,
	authUC authUseCase.UseCase,
	tokenParser middleware.TokenParser,
) *Server {
	srv := &http.Server{
		Addr: addr,
	}

	th := tasksHandler.New(taskUC)
	uh := usersHandler.New(userUC)
	ah := authHandler.New(authUC)

	r := configureRouter(th, uh, ah, tokenParser)

	srv.Handler = r

	return &Server{
		srv: srv,
	}
}

func configureRouter(
	th *tasksHandler.Handler,
	uh *usersHandler.Handler,
	ah *authHandler.Handler,
	parser middleware.TokenParser,
) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tasksGroup := r.Group("/tasks")
	tasksGroup.Use(middleware.AuthMiddleware(parser))
	tasksHandler.RegisterRoutes(tasksGroup, th)

	publicUsers := r.Group("/users")
	usersHandler.RegisterPublicRoutes(publicUsers, uh)

	privateUsers := r.Group("/users")
	privateUsers.Use(middleware.AuthMiddleware(parser))
	usersHandler.RegisterPrivateRoutes(privateUsers, uh)

	authGroup := r.Group("/auth")
	authHandler.RegisterRoutes(authGroup, ah)

	return r
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
