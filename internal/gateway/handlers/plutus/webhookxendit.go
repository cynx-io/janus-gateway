package plutus

import (
	"encoding/json"
	proto "github.com/cynx-io/janus-gateway/api/proto/gen/plutus"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type WebhookXenditHandler struct {
	client proto.WebhookXenditServiceClient
}

func NewWebhookXenditHandler() *WebhookXenditHandler {
	conn, err := grpc.NewClient(config.Config.Plutus.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Ananke gRPC server: " + err.Error())
	}
	client := proto.NewWebhookXenditServiceClient(conn)
	return &WebhookXenditHandler{client: client}
}

func (h *WebhookXenditHandler) HandlePaymentInvoice(w http.ResponseWriter, r *http.Request) {
	var req proto.HandlePaymentInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	req.WebhookKey = r.Header.Get("X_CALLBACK_TOKEN")

	resp, err := h.client.HandlePaymentInvoice(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
