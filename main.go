package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/albanybuipe96/bookrestapi/internal/configs"
	"github.com/albanybuipe96/bookrestapi/internal/handlers"
	"github.com/albanybuipe96/bookrestapi/internal/routers"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv(configs.PORT)

	if port == "" {
		log.Fatal("PORT not found in environment variables")
	}

	router := chi.NewRouter()
	router.Use(configs.Cors())

	// home route
	router.Get("/", handlers.Index)

	// v1 routes
	router.Mount("/v1", routers.V1Router())

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", port),
	}

	log.Printf("Listening on port: %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
		return
	}

}
