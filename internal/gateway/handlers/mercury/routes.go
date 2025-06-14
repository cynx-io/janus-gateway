package mercury

import "github.com/gorilla/mux"

func (h *CryptoHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/mercury.MercuryCryptoService").Subrouter()
	_ = privateRouter.PathPrefix("/mercury.MercuryCryptoService").Subrouter()

	public.HandleFunc("/SearchCoin", h.SearchCoin)
	public.HandleFunc("/GetCoinRisk", h.GetCoinRisk)
}
