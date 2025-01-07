package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/polupolu-dev/polupolu-backend/internal/domain/models"
	"github.com/polupolu-dev/polupolu-backend/internal/usecase"
)

type NewsHandler struct {
	usecase *usecase.NewsUsecase
}

func NewNewsHandler(u *usecase.NewsUsecase) *NewsHandler {
	return &NewsHandler{usecase: u}
}

func (h *NewsHandler) GetNewsDetail(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsID := c.Param("news_id")
		news, err := usecase.GetNewsDetail(newsID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, news)
	}
}

func (h *NewsHandler) GetCategoryNews(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Param("category")
		news, err := usecase.GetCategoryNews(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, news)
	}
}

func (h *NewsHandler) CreateNews(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newsRequest models.News
		if err := c.ShouldBindJSON(&newsRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		news, err := usecase.CreateNews(&newsRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, news)
	}
}

func (h *NewsHandler) GetAllNews(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		news, err := usecase.GetAllNews()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, news)
	}
}

func (h *NewsHandler) DeleteNews(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsID := c.Param("news_id")
		err := usecase.DeleteNews(newsID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"message": "News deleted successfully"})
	}
}

func (h *NewsHandler) UpdateNews(usecase *usecase.NewsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newsRequest models.News
		if err := c.ShouldBindJSON(&newsRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		news, err := usecase.UpdateNews(&newsRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, news)
	}
}
