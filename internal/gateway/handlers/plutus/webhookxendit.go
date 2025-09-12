package plutus

import (
	"encoding/json"
	"errors"
	core "github.com/cynx-io/cynx-core/proto/gen"
	proto "github.com/cynx-io/janus-gateway/api/proto/gen/plutus"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type WebhookXenditHandler struct {
	client proto.WebhookXenditServiceClient
}

// XenditPaymentInvoice represents the JSON structure from Xendit webhook
type XenditPaymentInvoice struct {
	PaidAt                 string `json:"paid_at"`
	ExternalID             string `json:"external_id"`
	UserID                 string `json:"user_id"`
	PaymentDestination     string `json:"payment_destination"`
	PaymentMethod          string `json:"payment_method"`
	Status                 string `json:"status"`
	MerchantName           string `json:"merchant_name"`
	PaymentChannel         string `json:"payment_channel"`
	BankCode               string `json:"bank_code"`
	Currency               string `json:"currency"`
	Description            string `json:"description"`
	PayerEmail             string `json:"payer_email"`
	ID                     string `json:"id"`
	Created                string `json:"created"`
	Updated                string `json:"updated"`
	FeesPaidAmount         int32  `json:"fees_paid_amount"`
	AdjustedReceivedAmount int32  `json:"adjusted_received_amount"`
	PaidAmount             int32  `json:"paid_amount"`
	Amount                 int32  `json:"amount"`
	IsHigh                 bool   `json:"is_high"`
}

// parseXenditTimestamp parses Xendit timestamp string to timestamppb.Timestamp
func parseXenditTimestamp(timeStr string) (*timestamppb.Timestamp, error) {
	if timeStr == "" {
		return nil, nil
	}

	// Try common timestamp formats that Xendit might use
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return timestamppb.New(t), nil
		}
	}

	return nil, nil // Return nil if parsing fails
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
	// First decode into Xendit structure
	var xenditReq XenditPaymentInvoice
	if err := json.NewDecoder(r.Body).Decode(&xenditReq); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Convert to protobuf structure
	req := &proto.HandlePaymentInvoiceRequest{
		Id:                     xenditReq.ID,
		ExternalId:             xenditReq.ExternalID,
		UserId:                 xenditReq.UserID,
		IsHigh:                 xenditReq.IsHigh,
		PaymentMethod:          xenditReq.PaymentMethod,
		Status:                 xenditReq.Status,
		MerchantName:           xenditReq.MerchantName,
		Amount:                 xenditReq.Amount,
		PaidAmount:             xenditReq.PaidAmount,
		BankCode:               xenditReq.BankCode,
		PayerEmail:             xenditReq.PayerEmail,
		Description:            xenditReq.Description,
		AdjustedReceivedAmount: xenditReq.AdjustedReceivedAmount,
		FeesPaidAmount:         xenditReq.FeesPaidAmount,
		Currency:               xenditReq.Currency,
		PaymentChannel:         xenditReq.PaymentChannel,
		PaymentDestination:     xenditReq.PaymentDestination,
		WebhookKey:             r.Header.Get("X-CALLBACK-TOKEN"),
		Base:                   &core.BaseRequest{},
	}

	// Parse timestamps
	if paidAt, err := parseXenditTimestamp(xenditReq.PaidAt); err == nil {
		req.PaidAt = paidAt
	}
	if updated, err := parseXenditTimestamp(xenditReq.Updated); err == nil {
		req.Updated = updated
	}
	if created, err := parseXenditTimestamp(xenditReq.Created); err == nil {
		req.Created = created
	}

	resp, err := h.client.HandlePaymentInvoice(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.Base.Status != core.Status_SUCCESS {
		http.Error(w, errors.New("got status"+resp.Base.Status.String()+"  from backend with code "+resp.Base.Code+" "+resp.Base.Desc).Error(), http.StatusInternalServerError)
		return
	}

	err = handlers.HandleResponse(w, resp)
	if err != nil {
		http.Error(w, "Failed to handle response", http.StatusInternalServerError)
		return
	}
}
