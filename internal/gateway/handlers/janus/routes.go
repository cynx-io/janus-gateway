package janus

import (
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/hermes"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayHandler struct {
	userClient pb.HermesUserServiceClient
}

func NewGatewayHandler() *GatewayHandler {
	conn, err := grpc.NewClient(config.Config.Hermes.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Hermes gRPC server: " + err.Error())
	}

	userClient := pb.NewHermesUserServiceClient(conn)
	return &GatewayHandler{userClient: userClient}
}

func (h *GatewayHandler) InjectRoutes(router *mux.Router) {

	auth0 := router.PathPrefix("/auth0").Subrouter()

	auth0.HandleFunc("/login", h.Auth0Login).Methods("GET")
	auth0.HandleFunc("/callback", h.Auth0CallbackLogin).Methods("GET")
	auth0.HandleFunc("/me", h.Auth0Me).Methods("GET")
	auth0.HandleFunc("/logout", h.Auth0Logout).Methods("GET", "POST")

}
