package ananke

import "github.com/gorilla/mux"

func (h *PreorderHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	_ = publicRouter.PathPrefix("/ananke.PreorderService").Subrouter()
	private := privateRouter.PathPrefix("/ananke.PreorderService").Subrouter()

	private.HandleFunc("/InitiatePreorder", h.InitiatePreorder)
	private.HandleFunc("/GetLatestCompletedOrPendingPreorder", h.GetLatestCompletedOrPendingPreorder)
}
