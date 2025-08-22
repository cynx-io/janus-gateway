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

func (h *CareerProfileHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/philyra.CareerProfileService").Subrouter()
	private := privateRouter.PathPrefix("/philyra.CareerProfileService").Subrouter()

	public.HandleFunc("/GetCareerProfile", h.GetCareerProfile)

	private.HandleFunc("/SyncCareerProfile", h.SyncCareerProfile)
	private.HandleFunc("/UpdatePersonalInfo", h.UpdatePersonalInfo)
	private.HandleFunc("/UpdateProfessionalInfo", h.UpdateProfessionalInfo)
	private.HandleFunc("/UpdateJobPreferences", h.UpdateJobPreferences)
	private.HandleFunc("/UpdateUserDocuments", h.UpdateUserDocuments)
	private.HandleFunc("/UpdateCustomResponses", h.UpdateCustomResponses)
}

func (h *AutoFillHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	private := privateRouter.PathPrefix("/philyra.AutoFillService").Subrouter()

	private.HandleFunc("/AnalyzeForm", h.AnalyzeForm)
}
