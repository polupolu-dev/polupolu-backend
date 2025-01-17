package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/polupolu-dev/polupolu-backend/internal/adapter/handlers"
	"github.com/polupolu-dev/polupolu-backend/internal/adapter/router"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/infrastructure/postgres"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
	"github.com/polupolu-dev/polupolu-backend/utils/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open("postgres", config.DB_DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	commentRepo := postgres.NewCommentRepository(dbConn)
	newsRepo := postgres.NewNewsRepository(dbConn)
	userRepo := postgres.NewUserRepository(dbConn)
	var llmService interfaces.LLMService

	commentUC := usecase.NewCommentUsecase(commentRepo, newsRepo, userRepo, llmService)
	newsUC := usecase.NewNewsUsecase(newsRepo, llmService)
	userUC := usecase.NewUserUsecase(userRepo)

	commentHandler := handlers.NewCommentHandler(commentUC)
	newsHandler := handlers.NewNewsHandler(newsUC)
	userHandler := handlers.NewUserHandler(userUC)

	r := router.NewRouter(commentHandler, newsHandler, userHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
