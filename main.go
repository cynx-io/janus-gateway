package main

import (
	"errors"
	"github.com/gorilla/mux"
	"janus/internal/gateway/handlers"
	"janus/internal/gateway/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Starting Janus API Gateway")

	// Create user handler
	userHandler, err := handlers.NewUserHandler()
	if err != nil {
		log.Fatalf("Failed to create user handler: %v", err)
	}

	// Create router
	publicRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	publicRouter.Use(middleware.CORSMiddleware)

	privateRouter := publicRouter.PathPrefix("").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	// Register routes
	publicRouter.HandleFunc("/user/check_username", userHandler.CheckUsername)
	publicRouter.HandleFunc("/user/register", userHandler.CreateUser)
	publicRouter.HandleFunc("/user/login", userHandler.ValidatePassword)

	privateRouter.HandleFunc("/user/profile", userHandler.GetUser)

	address := ":" + strconv.Itoa(middleware.GetConfig().App.Port)

	// Create server with middleware
	server := &http.Server{
		Addr:    address,
		Handler: publicRouter,
	}

	// Start server
	log.Printf("HTTP server listening on %s", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server: %v", err)
	}
}
