package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type CommentHandler struct {
	commentUseCase *usecase.CommentUsecase
}

func NewCommentHandler(cu *usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{commentUseCase: cu}
}

func (h *CommentHandler) GetCommentForNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	newsID := vars["news_id"]
	comments, err := h.commentUseCase.GetCommentForNews(r.Context(), newsID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func (h *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["comment_id"]
	comment, err := h.commentUseCase.GetComment(r.Context(), commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.commentUseCase.CreateComment(r.Context(), &comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/api/v1/comments/"+comment.ID)
	w.WriteHeader(http.StatusCreated)
}
