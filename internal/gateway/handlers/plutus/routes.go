package plutus

import "github.com/gorilla/mux"

func (h *WebhookXenditHandler) InjectRoutes(webhookRouter *mux.Router) {
	router := webhookRouter.PathPrefix("/plutus.WebhookXenditService").Subrouter()

	router.HandleFunc("/HandlePaymentInvoice", h.HandlePaymentInvoice)
}

func (h *PaymentHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/plutus.PaymentService").Subrouter()
	private := privateRouter.PathPrefix("/plutus.PaymentService").Subrouter()

	public.HandleFunc("/ListTokenPriceList", h.ListTokenPriceList)
	public.HandleFunc("/GetTokenPriceListById", h.GetTokenPriceListById)
	public.HandleFunc("/GetProductPriceListById", h.GetProductPriceListById)

	private.HandleFunc("/GetBalance", h.GetBalance)
	private.HandleFunc("/TopUpBalance", h.TopUpBalance)
	private.HandleFunc("/PurchaseProduct", h.PurchaseProduct)
}
