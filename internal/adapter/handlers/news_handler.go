package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type NewsHandler struct {
	newsUseCase *usecase.NewsUsecase
}

func NewNewsHandler(nu *usecase.NewsUsecase) *NewsHandler {
	return &NewsHandler{newsUseCase: nu}
}

func (h *NewsHandler) GetNewsDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["news_id"]
	news, err := h.newsUseCase.GetNewsDetail(r.Context(), newsID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) GetCategoryNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	news, err := h.newsUseCase.GetCategoryNews(r.Context(), category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) CreateNews(w http.ResponseWriter, r *http.Request) {
	var news models.News
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.newsUseCase.CreateNews(r.Context(), &news); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/api/v1/news/"+news.ID)
	w.WriteHeader(http.StatusCreated)
}

func (h *NewsHandler) GetAllNews(w http.ResponseWriter, r *http.Request) {
	news, err := h.newsUseCase.GetAllNews(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) DeleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["news_id"]
	err := h.newsUseCase.DeleteNews(r.Context(), newsID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *NewsHandler) UpdateNews(w http.ResponseWriter, r *http.Request) {
	var news models.News
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.newsUseCase.UpdateNews(r.Context(), &news); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/api/v1/news/"+news.ID)
	w.WriteHeader(http.StatusCreated)
}
