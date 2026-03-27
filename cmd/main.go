package main

import (
	"context"
	"log"
	"todos-api/internal/app"
	"todos-api/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	server "todos-api/internal/app/http"
	taskRepo "todos-api/internal/repository/postgres/tasks"
	userRepo "todos-api/internal/repository/postgres/users"
	taskUsecase "todos-api/internal/usecase/tasks"

	userUsecase "todos-api/internal/usecase/users"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer <token>

func main() {
	_ = godotenv.Load()

	cfg := config.MustLoad()

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	app.RunMigrations(cfg.DatabaseURL)

	tr := taskRepo.NewRepository(db)
	tuc := taskUsecase.New(tr)

	ur := userRepo.NewRepository(db)
	uuc := userUsecase.New(ur)

	srv := server.New(":"+cfg.HTTPPort, tuc, uuc)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
