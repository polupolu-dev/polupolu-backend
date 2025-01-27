package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/polupolu-dev/polupolu-backend/internal/adapter/handlers"
	"github.com/polupolu-dev/polupolu-backend/internal/adapter/router"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/interfaces"
	"github.com/polupolu-dev/polupolu-backend/internal/infrastructure/llm"
	"github.com/polupolu-dev/polupolu-backend/internal/infrastructure/postgres"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
	"github.com/polupolu-dev/polupolu-backend/utils/config"
	"github.com/polupolu-dev/polupolu-backend/utils/consts"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	var dbConn *sql.DB

	switch config.DBType {
	case consts.Cloudsql:
		var err error
		dbConn, err = config.Cloudsql()
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close()
	case consts.Postgres:
		var err error
		dbConn, err = config.Postgres()
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close()
	}

	commentRepo := postgres.NewCommentRepository(dbConn)
	newsRepo := postgres.NewNewsRepository(dbConn)
	userRepo := postgres.NewUserRepository(dbConn)

	llmService := &llm.NoopLLMService{}

	commentUC := usecase.NewCommentUsecase(
		commentRepo, newsRepo, userRepo, llmService)
	newsUC := usecase.NewNewsUsecase(newsRepo, llmService)
	userUC := usecase.NewUserUsecase(userRepo)

	commentHandler := handlers.NewCommentHandler(commentUC)
	newsHandler := handlers.NewNewsHandler(newsUC)
	userHandler := handlers.NewUserHandler(userUC)

	r := router.NewRouter(commentHandler, newsHandler, userHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
