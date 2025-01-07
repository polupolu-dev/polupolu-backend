package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type CommentsHandler struct {
	usecase *usecase.CommentsUsecase
}

func NewCommentsHandler(u *usecase.CommentsUsecase) *CommentsHandler {
	return &CommentsHandler{usecase: u}
}

func (h *CommentsHandler) GetCommentsForNews(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsID := c.Param("news_id")
		comments, err := usecase.GetCommentsForNews(newsID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}

func (h *CommentsHandler) GetComment(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")
		comment, err := usecase.GetComment(commentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, comment)
	}
}

func (h *CommentsHandler) GetUserComments(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		comments, err := usecase.GetUserComments(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}

func (h *CommentsHandler) CreateComment(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentRequest models.Comment
		if err := c.BindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := usecase.CreateComment(&commentRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, comment)
	}
}

func (h *CommentsHandler) CreateReply(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentRequest models.Comment
		if err := c.BindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := usecase.CreateReply(&commentRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, comment)
	}
}

func (h *CommentsHandler) DeleteComment(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")
		err := usecase.DeleteComment(commentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "comment deleted"})
	}
}

func (h *CommentsHandler) UpdateComment(usecase usecase.CommentsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentRequest models.Comment
		if err := c.BindJSON(&commentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		comment, err := usecase.UpdateComment(&commentRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, comment)
	}
}
