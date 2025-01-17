package router

import (
	"github.com/gorilla/mux"
	"github.com/polupolu-dev/polupolu-backend/internal/adapter/handlers"
)

func NewRouter(
	commentHandler *handlers.CommentHandler,
	newsHandler *handlers.NewsHandler,
	userHandler *handlers.UserHandler,
) *mux.Router {
	r := mux.NewRouter()

	// Comments
	r.HandleFunc("/api/v1/news/{news_id}/comments", commentHandler.GetCommentsForNews).Methods("GET")
	r.HandleFunc("/api/v1/comments/{comment_id}", commentHandler.GetComment).Methods("GET")
	r.HandleFunc("/api/v1/users/{user_id}/comments", commentHandler.GetUserComments).Methods("GET")
	r.HandleFunc("/api/v1/news/{news_id}/comments", commentHandler.CreateComment).Methods("POST")
	r.HandleFunc("/api/v1/comments/{comment_id}/replies", commentHandler.CreateReply).Methods("POST")
	r.HandleFunc("/api/v1/comments/{comment_id}", commentHandler.DeleteComment).Methods("DELETE")
	r.HandleFunc("/api/v1/comments/{comment_id}", commentHandler.UpdateComment).Methods("PUT")
	// News
	r.HandleFunc("/api/v1/news/{news_id}", newsHandler.GetNewsDetail).Methods("GET")
	r.HandleFunc("/api/v1/news/categories/{category}", newsHandler.GetCategoryNews).Methods("GET")
	r.HandleFunc("/api/v1/news", newsHandler.CreateNews).Methods("POST")
	r.HandleFunc("/api/v1/news", newsHandler.GetAllNews).Methods("GET")
	r.HandleFunc("/api/v1/news/{news_id}", newsHandler.DeleteNews).Methods("DELETE")
	r.HandleFunc("/api/v1/news/{news_id}", newsHandler.UpdateNews).Methods("PUT")
	// Users
	r.HandleFunc("/api/v1/users/{user_id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{user_id}", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/v1/users/{user_id}", userHandler.UpdateUser).Methods("PUT")
	return r
}
