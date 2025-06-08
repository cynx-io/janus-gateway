package handlers

import "github.com/gorilla/mux"

func (h *UserHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {

	public := publicRouter.PathPrefix("/user").Subrouter()
	private := privateRouter.PathPrefix("/user").Subrouter()

	public.HandleFunc("/check_username", h.CheckUsername)
	public.HandleFunc("/register", h.CreateUser)
	public.HandleFunc("/login", h.ValidatePassword)

	private.HandleFunc("/profile", h.GetUser)
}

func (h *CryptoHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/crypto").Subrouter()
	_ = privateRouter.PathPrefix("/crypto").Subrouter()

	public.HandleFunc("/eth/search", h.SearchCoin)
	public.HandleFunc("/eth/risk", h.GetCoinRisk)
}
