package plutus

import "github.com/gorilla/mux"

func (h *WebhookXenditHandler) InjectRoutes(webhookRouter *mux.Router) {
	router := webhookRouter.PathPrefix("/plutus.WebhookXenditService").Subrouter()

	router.HandleFunc("/HandlePaymentInvoice", h.HandlePaymentInvoice)
}
