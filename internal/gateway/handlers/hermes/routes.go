package hermes

import (
	"github.com/gorilla/mux"
)

func (h *UserHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {

	public := publicRouter.PathPrefix("/hermes.HermesUserService").Subrouter()
	private := privateRouter.PathPrefix("/hermes.HermesUserService").Subrouter()

	public.HandleFunc("/CheckUsername", h.CheckUsername)
	public.HandleFunc("/CreateUser", h.CreateUser)
	public.HandleFunc("/ValidatePassword", h.ValidatePassword)

	private.HandleFunc("/GetUser", h.GetUser)
}
