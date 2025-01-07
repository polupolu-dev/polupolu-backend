package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type UsersHandler struct {
	usecase *usecase.UsersUsecase
}

func NewUsersHandler(u *usecase.UsersUsecase) *UsersHandler {
	return &UsersHandler{usecase: u}
}

func (h *UsersHandler) GetUsers(usecase *usecase.UsersUsecase) gin.HandlerFunc {
	return func(u *gin.Context) {
		userID := u.Param("user_id")
		users, err := h.usecase.GetUser(userID)
		if err != nil {
			u.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		u.JSON(http.StatusOK, users)
	}
}

func (h *UsersHandler) CreateUser(usecase *usecase.UsersUsecase) gin.HandlerFunc {
	return func(u *gin.Context) {
		var user models.User
		if err := u.BindJSON(&user); err != nil {
			u.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdUser, err := h.usecase.CreateUser(&user)
		if err != nil {
			u.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		u.JSON(http.StatusCreated, createdUser)
	}
}

func (h *UsersHandler) DeleteUser(usecase *usecase.UsersUsecase) gin.HandlerFunc {
	return func(u *gin.Context) {
		userID := u.Param("user_id")
		if err := h.usecase.DeleteUser(userID); err != nil {
			u.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		u.Status(http.StatusNoContent)
	}
}

func (h *UsersHandler) UpdateUser(usecase *usecase.UsersUsecase) gin.HandlerFunc {
	return func(u *gin.Context) {
		var user models.User
		if err := u.BindJSON(&user); err != nil {
			u.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedUser, err := h.usecase.UpdateUser(&user)
		if err != nil {
			u.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		u.JSON(http.StatusOK, updatedUser)
	}
}
