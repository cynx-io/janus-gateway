package athena

import "github.com/gorilla/mux"

func (h *IdeaHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/athena.IdeaService").Subrouter()
	_ = privateRouter.PathPrefix("/athena.IdeaService").Subrouter()

	public.HandleFunc("/ValidateIdea", h.ValidateIdea)
}
