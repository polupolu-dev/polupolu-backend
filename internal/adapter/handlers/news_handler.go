package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// ... 他のハンドラーメソッドも同様に実装
