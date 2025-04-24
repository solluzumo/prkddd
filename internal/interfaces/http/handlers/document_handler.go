package handlers

import (
	"context"
	"math"
	"net/http"
	"prk/internal/application/document"
	documentD "prk/internal/domain/document"
	"prk/internal/interfaces/http/middleware"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type PaginatedResponse struct {
	Data       []*documentD.Document
	Total      int64
	Page       int
	PerPage    int
	TotalPages int
}

type DocumentFilterRequest struct {
	DocTypeID         *int   `json:"doc_type_id" schema:"doc_type_id"` // schema для github.com/gorilla/schema
	UpdateRegularly   *bool  `json:"update" schema:"update"`
	NeedsExpertReview *bool  `json:"review" schema:"review"`
	JournalCategory   string `json:"journal_category" schema:"journal_category"`
	Status            string `json:"status" schema:"status"`
}

type DocumentHandler struct {
	service *document.DocumentService
}

func NewDocumentHandler(service *document.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

func (h *DocumentHandler) DocumentCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		documentID := chi.URLParam(r, "documentID")
		if documentID == "" {
			return
		}

		doc, err := h.service.FindDocumentById(documentID)
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), "document", doc)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *DocumentHandler) GetDocuments(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	//Пагинация
	page, _ := strconv.Atoi(queryParams.Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 20
	}
	//Сортировка
	sortField := queryParams.Get("sort")
	if sortField == "" {
		sortField = "author"
	}
	sortOrder := queryParams.Get("order")
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}
	//Настраиваем фильтры
	filters := make(map[string]interface{})
	if typeID := queryParams.Get("doc_type_id"); typeID != "" {
		if id, err := strconv.Atoi(typeID); err == nil {
			filters["document_type"] = id
		} else {
			return
		}
	}

	if update := queryParams.Get("update"); update != "" {
		filters["updated_regularly"] = update
	}

	if review := queryParams.Get("review"); review != "" {
		filters["expert_review"] = review
	}

	if category := queryParams.Get("journal_category"); category != "" {
		filters["journal_category"] = category
	}

	if status := queryParams.Get("status"); status != "" {
		filters["status"] = status
	}
	//Формируем дто-шки
	form := document.ListDoucmentDTO{
		Page:      page,
		Limit:     limit,
		SortField: sortField,
		SortOrder: sortOrder,
		Filters:   filters,
	}
	documents, total, err := h.service.FindDocuments(r.Context(), form)
	if err != nil {
		return
	}
	response := PaginatedResponse{
		Data:       documents,
		Total:      total,
		Page:       page,
		PerPage:    limit,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}
	render.JSON(w, r, response)
}

func (h *DocumentHandler) AddDocument(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.TokenKey)
	if token == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	defer r.Body.Close()
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		if err == http.ErrNotMultipart {
			http.Error(w, "Request must be multipart/form-data", http.StatusBadRequest)
		} else if strings.Contains(err.Error(), "request body too large") {
			http.Error(w, "File too large (max 10MB)", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
		}
		return
	}

	mainFileHeader := r.MultipartForm.File["mainFile"]
	if len(mainFileHeader) == 0 {
		http.Error(w, "Main file required", http.StatusBadRequest)
		return
	}
	mainFile, err := mainFileHeader[0].Open()
	if err != nil {
		http.Error(w, "Can't open main file", http.StatusInternalServerError)
		return
	}

	form := document.CreateDocumentDTO{
		DocumentTypeID:    r.FormValue("document_type"),
		Title:             r.FormValue("title"),
		Date:              r.FormValue("date"),
		UpdatedRegularly:  r.FormValue("updated_regularly") == "true",
		ExpertReview:      r.FormValue("expert_review") == "true",
		JournalCategoryID: r.FormValue("journal_category"),
		Source:            r.FormValue("source"),
		MainFile:          mainFile,
		MainFileName:      mainFileHeader[0].Filename,
		AdditionalFiles:   r.MultipartForm.File["additionFiles"],
	}

	err = h.service.CreateDocument(r.Context(), token.(string), form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Document added successfully"}`))
}
