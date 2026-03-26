package main

import (
	"context"
	"log"
	"todos-api/internal/app"
	"todos-api/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	server "todos-api/internal/app/http"
	repo "todos-api/internal/repository/postgres/tasks"
	usecase "todos-api/internal/usecase/tasks"
)

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

	taskRepo := repo.NewRepository(db)
	taskUC := usecase.New(taskRepo)

	srv := server.New(":"+cfg.HTTPPort, taskUC)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
