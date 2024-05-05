package routers

import (
	"database/sql"
	"log"
	"os"

	"github.com/albanybuipe96/bookrestapi/internal/configs"
	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/albanybuipe96/bookrestapi/internal/handlers"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func V1Router() *chi.Mux {

	dbUrl := os.Getenv(configs.DB_URL)

	if dbUrl == "" {
		log.Fatal("DB_URL not found in environment variables")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatalf("error connecting to db: %v", err.Error())
	}

	dbConfig := handlers.DbConfig{
		DB: database.New(conn),
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/", handlers.Index)
	v1Router.Get("/error", handlers.Error)

	// users routes
	v1Router.Post("/users", dbConfig.CreateUser)
	v1Router.Get("/users", dbConfig.AuthMiddleware(dbConfig.GetUserByAPIKey))

	// books routes
	v1Router.Post("/books", dbConfig.AuthMiddleware(dbConfig.AddBook))
	v1Router.Get("/books", dbConfig.GetBooks)

	return v1Router
}
