package philyra

import "github.com/gorilla/mux"

func (h *ResumeHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/philyra.ResumeService").Subrouter()
	private := privateRouter.PathPrefix("/philyra.ResumeService").Subrouter()

	public.HandleFunc("/GetResume", h.GetResume)
	public.HandleFunc("/ListResumes", h.ListResumes)
	public.HandleFunc("/GenerateResume", h.GenerateResume)

	private.HandleFunc("/CreateResume", h.CreateResume)
	private.HandleFunc("/UpdateResume", h.UpdateResume)
	private.HandleFunc("/DeleteResume", h.DeleteResume)
}
