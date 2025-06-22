package plato

import "github.com/gorilla/mux"

func (h *AnswerHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/plato.PlatoAnswerService").Subrouter()
	private := privateRouter.PathPrefix("/plato.PlatoAnswerService").Subrouter()

	public.HandleFunc("/SearchAnswers", h.SearchAnswers)
	public.HandleFunc("/ListDetailAnswersByTopicId", h.ListDetailAnswersByTopicId)

	private.HandleFunc("/GetAnswerById", h.GetAnswerById)
	private.HandleFunc("/GetDetailAnswerById", h.GetDetailAnswerById)
	private.HandleFunc("/ListAnswersByTopicId", h.ListAnswersByTopicId)

	private.HandleFunc("/InsertAnswer", h.InsertAnswer)
	private.HandleFunc("/UpdateAnswer", h.UpdateAnswer)
	private.HandleFunc("/DeleteAnswer", h.DeleteAnswer)
}

func (h *AnswerCategoryHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	_ = publicRouter.PathPrefix("/plato.PlatoAnswerCategoryService").Subrouter()
	private := privateRouter.PathPrefix("/plato.PlatoAnswerCategoryService").Subrouter()

	private.HandleFunc("/GetAnswerCategoryById", h.GetAnswerCategoryById)
	private.HandleFunc("/ListAnswerCategoriesByAnswerId", h.ListAnswerCategoriesByAnswerId)

	private.HandleFunc("/InsertAnswerCategory", h.InsertAnswerCategory)
	private.HandleFunc("/UpdateAnswerCategory", h.UpdateAnswerCategory)
	private.HandleFunc("/DeleteAnswerCategory", h.DeleteAnswerCategory)
}

func (h *DailyGameHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/plato.PlatoDailyGameService").Subrouter()
	private := privateRouter.PathPrefix("/plato.PlatoDailyGameService").Subrouter()

	public.HandleFunc("/GetModeDailyGameById", h.GetModeDailyGameById)
	public.HandleFunc("/GetPublicDailyGame", h.GetPublicDailyGame)
	public.HandleFunc("/AttemptAnswer", h.AttemptAnswer)

	private.HandleFunc("/GetDetailDailyGameById", h.GetDetailDailyGameById)
}

func (h *ModeHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/plato.PlatoModeService").Subrouter()
	private := privateRouter.PathPrefix("/plato.PlatoModeService").Subrouter()

	public.HandleFunc("/ListModesByTopicId", h.ListModesByTopicId)

	private.HandleFunc("/GetModeById", h.GetModeById)
	private.HandleFunc("/InsertMode", h.InsertMode)
	private.HandleFunc("/UpdateMode", h.UpdateMode)
	private.HandleFunc("/DeleteMode", h.DeleteMode)
}

func (h *TopicHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/plato.PlatoTopicService").Subrouter()
	private := privateRouter.PathPrefix("/plato.PlatoTopicService").Subrouter()

	public.HandleFunc("/PaginateTopic", h.PaginateTopic)
	public.HandleFunc("/GetTopicBySlug", h.GetTopicBySlug)
	public.HandleFunc("/GetTopicById", h.GetTopicById)

	private.HandleFunc("/InsertTopic", h.InsertTopic)
	private.HandleFunc("/UpdateTopic", h.UpdateTopic)
	private.HandleFunc("/DeleteTopic", h.DeleteTopic)
	private.HandleFunc("/ListTopicsByUserId", h.ListTopicsByUserId)
}
