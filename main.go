package main

import (
	"errors"
	"github.com/gorilla/mux"
	"janus/internal/dependencies/config"
	"janus/internal/dependencies/logger"
	"janus/internal/gateway/handlers"
	"janus/internal/gateway/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Starting Janus API Gateway")

	// Load configuration
	logger.Init()
	config.Init()

	// Create user handler
	userHandler := handlers.NewUserHandler()
	cryptoHandler := handlers.NewCryptoHandler()

	// Create router
	publicRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	publicRouter.Use(middleware.CORSMiddleware)

	privateRouter := publicRouter.PathPrefix("").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	// Inject routes
	userHandler.InjectRoutes(publicRouter, privateRouter)
	cryptoHandler.InjectRoutes(publicRouter, privateRouter)

	address := ":" + strconv.Itoa(config.Config.App.Port)

	// Create server with middleware
	server := &http.Server{
		Addr:    address,
		Handler: publicRouter,
	}

	// Start server
	logger.Info("HTTP server listening on ", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("Failed to start server: " + err.Error())
	}
}
