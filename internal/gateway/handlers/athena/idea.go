package athena

import (
	"encoding/json"
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/athena"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type IdeaHandler struct {
	client pb.IdeaServiceClient
}

func NewIdeaHandler() *IdeaHandler {
	conn, err := grpc.NewClient(config.Config.Athena.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Athena gRPC server: " + err.Error())
	}
	client := pb.NewIdeaServiceClient(conn)
	return &IdeaHandler{client: client}
}

func (h *IdeaHandler) GenerateIdeas(w http.ResponseWriter, r *http.Request) {
	var req pb.GenerateIdeasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GenerateIdeas(r.Context(), &req)
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

func (h *IdeaHandler) InitiateIdea(w http.ResponseWriter, r *http.Request) {
	var req pb.InitiateIdeaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.InitiateIdea(r.Context(), &req)
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

func (h *IdeaHandler) GetValidationScore(w http.ResponseWriter, r *http.Request) {
	var req pb.GetValidationScoreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetValidationScore(r.Context(), &req)
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

func (h *IdeaHandler) GetTrendAnalysis(w http.ResponseWriter, r *http.Request) {
	var req pb.GetTrendAnalysisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetTrendAnalysis(r.Context(), &req)
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

func (h *IdeaHandler) GetMarketData(w http.ResponseWriter, r *http.Request) {
	var req pb.GetMarketDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetMarketData(r.Context(), &req)
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

func (h *IdeaHandler) GetCommunitySignals(w http.ResponseWriter, r *http.Request) {
	var req pb.GetCommunitySignalsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetCommunitySignals(r.Context(), &req)
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

func (h *IdeaHandler) GetCompetitionAnalysis(w http.ResponseWriter, r *http.Request) {
	var req pb.GetCompetitionAnalysisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetCompetitionAnalysis(r.Context(), &req)
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

func (h *IdeaHandler) GetProblemValidation(w http.ResponseWriter, r *http.Request) {
	var req pb.GetProblemValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetProblemValidation(r.Context(), &req)
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

func (h *IdeaHandler) GetBootstrapMetrics(w http.ResponseWriter, r *http.Request) {
	var req pb.GetBootstrapMetricsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetBootstrapMetrics(r.Context(), &req)
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

func (h *IdeaHandler) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	var req pb.GetRecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetRecommendation(r.Context(), &req)
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

func (h *IdeaHandler) GetCompetitionChart(w http.ResponseWriter, r *http.Request) {
	var req pb.GetCompetitionChartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetCompetitionChart(r.Context(), &req)
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

func (h *IdeaHandler) GetEngagementChart(w http.ResponseWriter, r *http.Request) {
	var req pb.GetEngagementChartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetEngagementChart(r.Context(), &req)
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
func (h *IdeaHandler) GetTrendChart(w http.ResponseWriter, r *http.Request) {
	var req pb.GetTrendChartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetTrendChart(r.Context(), &req)
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
func (h *IdeaHandler) GetMarketDemandChart(w http.ResponseWriter, r *http.Request) {
	var req pb.GetMarketDemandChartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetMarketDemandChart(r.Context(), &req)
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
func (h *IdeaHandler) GetRedditInsights(w http.ResponseWriter, r *http.Request) {
	var req pb.GetRedditInsightsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetRedditInsights(r.Context(), &req)
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

func (h *IdeaHandler) GetInputAnalysis(w http.ResponseWriter, r *http.Request) {
	var req pb.GetInputAnalysisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetInputAnalysis(r.Context(), &req)
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

func (h *IdeaHandler) GetGithubInsights(w http.ResponseWriter, r *http.Request) {
	var req pb.GetGithubInsightsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetGithubInsights(r.Context(), &req)
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

func (h *IdeaHandler) GetHackerNewsInsights(w http.ResponseWriter, r *http.Request) {
	var req pb.GetHackerNewsInsightsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetHackerNewsInsights(r.Context(), &req)
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
