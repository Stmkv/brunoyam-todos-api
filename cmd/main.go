package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"todos-api/internal/app"
	"todos-api/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	server "todos-api/internal/app/http"
	jsonTaskRepo "todos-api/internal/repository/json/tasks"
	postgresTaskRepo "todos-api/internal/repository/postgres/tasks"
	postgresUserRepo "todos-api/internal/repository/postgres/users"
	authUsecase "todos-api/internal/usecase/auth"
	taskUsecase "todos-api/internal/usecase/tasks"

	userUsecase "todos-api/internal/usecase/users"

	"todos-api/internal/lib/hasher"
	"todos-api/internal/lib/jwt"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer <token>

func main() {
	ctx := context.Background()
	_ = godotenv.Load()

	cfg := config.MustLoad()
	app.RunMigrations(cfg.DatabaseURL)
	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := waitForDB(ctx, db); err != nil {
		log.Fatal(err)
	}

	var tr taskUsecase.Repository
	var ur userUsecase.Repository

	errPgConnection := db.Ping(context.Background())
	if errPgConnection != nil {
		// Repository json
		tr = jsonTaskRepo.NewRepository(cfg.FilePathForSaveTasks)
	} else {
		// Repository postgres
		tr = postgresTaskRepo.NewRepository(db)
		ur = postgresUserRepo.NewRepository(db)
	}

	bcryptHasher := hasher.New()
	jwtManager := jwt.New(cfg.JWTSecret)

	// Usecase
	tuc := taskUsecase.New(tr)
	uuc := userUsecase.New(ur, bcryptHasher)
	auc := authUsecase.New(ur, bcryptHasher, jwtManager)

	srv := server.New(":"+cfg.HTTPPort, tuc, uuc, auc, jwtManager)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func waitForDB(ctx context.Context, db *pgxpool.Pool) error {
	maxAttempts := 5
	baseDelay := time.Second * 2

	for i := 0; i < maxAttempts; i++ {
		err := db.Ping(ctx)
		if err == nil {
			return nil
		}
		log.Printf(
			"Waiting for database to be ready... attempt #%d, max attempts %d",
			i,
			maxAttempts,
		)
		time.Sleep(baseDelay)
		baseDelay += 2
	}
	return fmt.Errorf("timed out waiting for database to become available")
}
