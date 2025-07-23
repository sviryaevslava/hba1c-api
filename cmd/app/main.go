package main

import (
	"go-api-practice/config"
	"go-api-practice/internal/handler"
	"go-api-practice/internal/middleware"
	"go-api-practice/pkg/logger"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	logger.Init()

	mux := http.NewServeMux()
	mux.Handle("/predict", middleware.AuthMiddleware(http.HandlerFunc(handler.PredictHandler)))

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
