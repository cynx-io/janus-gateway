package athena

import "github.com/gorilla/mux"

func (h *IdeaHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/athena.IdeaService").Subrouter()
	_ = privateRouter.PathPrefix("/athena.IdeaService").Subrouter()

	public.HandleFunc("/GenerateIdeas", h.GenerateIdeas)
	public.HandleFunc("/InitiateIdea", h.InitiateIdea)
	public.HandleFunc("/GetValidationScore", h.GetValidationScore)
	public.HandleFunc("/GetTrendAnalysis", h.GetTrendAnalysis)
	public.HandleFunc("/GetMarketData", h.GetMarketData)
	public.HandleFunc("/GetCommunitySignals", h.GetCommunitySignals)
	public.HandleFunc("/GetCompetitionAnalysis", h.GetCompetitionAnalysis)
	public.HandleFunc("/GetProblemValidation", h.GetProblemValidation)
	public.HandleFunc("/GetBootstrapMetrics", h.GetBootstrapMetrics)
	public.HandleFunc("/GetRecommendation", h.GetRecommendation)
	public.HandleFunc("/GetCompetitionChart", h.GetCompetitionChart)
	public.HandleFunc("/GetEngagementChart", h.GetEngagementChart)
	public.HandleFunc("/GetTrendChart", h.GetTrendChart)
	public.HandleFunc("/GetMarketDemandChart", h.GetMarketDemandChart)
	public.HandleFunc("/GetRedditInsights", h.GetRedditInsights)
	public.HandleFunc("/GetHackerNewsInsights", h.GetHackerNewsInsights)
	public.HandleFunc("/GetGithubInsights", h.GetGithubInsights)
	public.HandleFunc("/GetInputAnalysis", h.GetInputAnalysis)
}
