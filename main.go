package main

import (
	"errors"
	"github.com/cynxees/janus-gateway/internal/dependencies/config"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/hermes"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/mercury"
	"github.com/cynxees/janus-gateway/internal/gateway/handlers/plato"
	"github.com/cynxees/janus-gateway/internal/gateway/middleware"
	"github.com/elastic/go-elasticsearch"
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
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://152.53.169.236:32200/"},
	})
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
	publicRouter := mux.NewRouter().PathPrefix("").Subrouter()
	publicRouter.Use(middleware.CORSMiddleware)
	publicRouter.Use(middleware.PublicAuthMiddleware)
	publicRouter.Use(logMiddleware.Handler)

	privateRouter := publicRouter.PathPrefix("").Subrouter()
	privateRouter.Use(middleware.PrivateAuthMiddleware)

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
		Handler: publicRouter,
	}

	// Start server
	logger.Info("HTTP server listening on ", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("Failed to start server: " + err.Error())
	}
}
