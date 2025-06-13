package mercury

import "github.com/gorilla/mux"

func (h *CryptoHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/crypto").Subrouter()
	_ = privateRouter.PathPrefix("/crypto").Subrouter()

	public.HandleFunc("/eth/search", h.SearchCoin)
	public.HandleFunc("/eth/risk", h.GetCoinRisk)
}
