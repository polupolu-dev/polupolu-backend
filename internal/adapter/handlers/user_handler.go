package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UsersUsecase
}

func NewUserHandler(uu *usecase.UsersUsecase) *UserHandler {
	return &UserHandler{userUseCase: uu}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	user, err := h.userUseCase.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// ... 他のハンドラーメソッドも同様に実装
