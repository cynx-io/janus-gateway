package main

import (
	"errors"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/dependencies/elastic"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/hermes"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/mercury"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/plato"
	"github.com/cynxees/janus-gateway/internal/gateway/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Starting Janus API Gateway")

	// Load configuration
	logger.Init()
	config.Init()

	// Load Dependencies
	es, err := elastic.NewClient()
	if err != nil {
		panic("Failed to create Elasticsearch client: " + err.Error())
	}

	logMiddleware := middleware.LoggingMiddleware{
		ElasticClient: es,
	}

	// Create user handler
	userHandler := hermes.NewUserHandler()
	cryptoHandler := mercury.NewCryptoHandler()

	platoAnswerHandler := plato.NewAnswerHandler()
	platoAnswerCategoryHandler := plato.NewAnswerCategoryHandler()
	platoDailyGameHandler := plato.NewDailyGameHandler()
	platoModeHandler := plato.NewModeHandler()
	platoTopicHandler := plato.NewTopicHandler()

	// Create router
	root := mux.NewRouter()
	root.Use(middleware.CORSMiddleware)
	root.Use(middleware.BaseRequestHandler)

	publicRouter := root.PathPrefix("").Subrouter()
	publicRouter.Use(
		middleware.PublicAuthMiddleware,
		logMiddleware.RequestHandler,
	)
	publicRouter.Use(logMiddleware.ResponseHandler)

	privateRouter := root.PathPrefix("/").Subrouter()
	privateRouter.Use(
		middleware.PrivateAuthMiddleware,
		logMiddleware.RequestHandler,
	)
	privateRouter.Use(logMiddleware.ResponseHandler)

	// Inject routes
	userHandler.InjectRoutes(publicRouter, privateRouter)
	cryptoHandler.InjectRoutes(publicRouter, privateRouter)
	platoAnswerHandler.InjectRoutes(publicRouter, privateRouter)
	platoAnswerCategoryHandler.InjectRoutes(publicRouter, privateRouter)
	platoDailyGameHandler.InjectRoutes(publicRouter, privateRouter)
	platoModeHandler.InjectRoutes(publicRouter, privateRouter)
	platoTopicHandler.InjectRoutes(publicRouter, privateRouter)

	address := ":" + strconv.Itoa(config.Config.App.Port)

	// Create server with middleware
	server := &http.Server{
		Addr:    address,
		Handler: root,
	}

	// Start server
	logger.Info("HTTP server listening on ", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("Failed to start server: " + err.Error())
	}
}
