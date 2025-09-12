package pheme

import (
	"encoding/json"
	"github.com/cynx-io/cynx-core/src/context"
	pb "github.com/cynx-io/janus-gateway/api/proto/gen/pheme"
	"github.com/cynx-io/janus-gateway/internal/dependencies/config"
	"github.com/cynx-io/janus-gateway/internal/gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	client pb.BlogServiceClient
}

func NewBlogHandler() *BlogHandler {
	conn, err := grpc.NewClient(config.Config.Pheme.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Pheme gRPC server: " + err.Error())
	}
	client := pb.NewBlogServiceClient(conn)
	return &BlogHandler{client: client}
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateBlog(r.Context(), &req)
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

func (h *BlogHandler) PaginateBlogs(w http.ResponseWriter, r *http.Request) {
	var req pb.PaginateBlogsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.PaginateBlogs(r.Context(), &req)
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

func (h *BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {
	var req pb.GetBlogByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetBlogById(r.Context(), &req)
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

func (h *BlogHandler) GetBlogByTitle(w http.ResponseWriter, r *http.Request) {
	var req pb.GetBlogByTitleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.client.GetBlogByTitle(r.Context(), &req)
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

func (h *BlogHandler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "Failed to close file", http.StatusInternalServerError)
			return
		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	filename := r.FormValue("filename")
	if filename == "" {
		filename = header.Filename
	}

	mediaContentTypeStr := r.FormValue("media_content_type")
	mediaContentType, err := strconv.Atoi(mediaContentTypeStr)
	if err != nil {
		http.Error(w, "Invalid media content type", http.StatusBadRequest)
		return
	}

	baseReq := context.GetBaseRequest(r.Context())

	req := pb.UploadMediaRequest{
		Base:             baseReq,
		File:             fileBytes,
		Filename:         filename,
		MediaContentType: pb.MediaContentType(mediaContentType),
	}

	resp, err := h.client.UploadMedia(r.Context(), &req)
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
