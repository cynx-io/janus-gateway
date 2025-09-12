package pheme

import "github.com/gorilla/mux"

func (h *BlogHandler) InjectRoutes(publicRouter *mux.Router, privateRouter *mux.Router) {
	public := publicRouter.PathPrefix("/pheme.BlogService").Subrouter()
	private := privateRouter.PathPrefix("/pheme.BlogService").Subrouter()

	public.HandleFunc("/GetBlogById", h.GetBlogById)
	public.HandleFunc("/GetBlogByTitle", h.GetBlogByTitle)
	public.HandleFunc("/PaginateBlogs", h.PaginateBlogs)

	private.HandleFunc("/CreateBlog", h.CreateBlog)
	private.HandleFunc("/UploadMedia", h.UploadMedia)
}
