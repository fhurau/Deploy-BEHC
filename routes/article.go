package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.HandleFunc("/articles", h.FindArticles).Methods("GET")
	r.HandleFunc("/article/{id}", h.GetArticle).Methods("GET")
	r.HandleFunc("/article", middleware.Auth(middleware.UploadFile(h.CreateArticle))).Methods("POST")
	r.HandleFunc("/article/{id}", middleware.Auth(middleware.UploadFile(h.UpdateArticle))).Methods("PATCH")
	r.HandleFunc("/article/{id}", middleware.Auth(h.DeleteArticle)).Methods("DELETE")
}
