package janus

import (
	"github.com/gorilla/mux"
)

type GatewayHandler struct {
}

func (h *GatewayHandler) InjectRoutes(router *mux.Router) {

	auth0 := router.PathPrefix("/auth0").Subrouter()

	auth0.HandleFunc("/login", h.Auth0Login).Methods("GET")
	auth0.HandleFunc("/callback", h.Auth0CallbackLogin).Methods("GET")

}
