package main

import (
	"context"
	"errors"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/janus-gateway/internal/dependencies/auth0"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers/hermes"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers/janus"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers/mercury"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers/philyra"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers/plato"
	"github.com/cynx-io/janus-gateway/internal/gateway/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.Println("Starting Janus API Gateway")

	// Load configuration
	config.Init()
	auth0.Init()

	logLevel, err := logrus.ParseLevel(config.Config.Elastic.Level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logger.Init(logger.LoggerConfig{
		Level:            logLevel,
		ElasticsearchURL: []string{config.Config.Elastic.Url},
		ServiceName:      "janus-gateway",
	})

	janusHandler := janus.GatewayHandler{}
	userHandler := hermes.NewUserHandler()
	cryptoHandler := mercury.NewCryptoHandler()
	resumeHandler := philyra.NewResumeHandler()

	platoAnswerHandler := plato.NewAnswerHandler()
	platoAnswerCategoryHandler := plato.NewAnswerCategoryHandler()
	platoDailyGameHandler := plato.NewDailyGameHandler()
	platoModeHandler := plato.NewModeHandler()
	platoTopicHandler := plato.NewTopicHandler()

	// Create router
	root := mux.NewRouter()
	janusHandler.InjectRoutes(root)

	root.Use(middleware.CORSMiddleware)

	publicRouter := root.PathPrefix("").Subrouter()
	publicRouter.Use(
		middleware.PublicAuthMiddleware,
		middleware.BaseRequestHandler,
		middleware.LogRequestHandler,
	)
	publicRouter.Use(middleware.LogResponseHandler)

	privateRouter := root.PathPrefix("/").Subrouter()
	privateRouter.Use(
		middleware.PrivateAuthMiddleware,
		middleware.BaseRequestHandler,
		middleware.LogRequestHandler,
	)
	privateRouter.Use(middleware.LogResponseHandler)

	// Inject routes
	userHandler.InjectRoutes(publicRouter, privateRouter)
	cryptoHandler.InjectRoutes(publicRouter, privateRouter)
	resumeHandler.InjectRoutes(publicRouter, privateRouter)
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
	logger.Info(context.Background(), "HTTP server listening on ", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("Failed to start server: " + err.Error())
	}
}
