package handlers

import (
	articledto "backend/dto/article"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	// "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

var path_file = "http://localhost:5000/uploads/"

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) FindArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := h.ArticleRepository.FindArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range articles {
		articles[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var article models.Article
	article, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article.Image = path_file + article.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticle(article)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "hellocorona"})

	if err != nil {
		fmt.Println(err.Error())
	}

	request := articledto.ArticleRequest{
		Title:  r.FormValue("title"),
		UserID: userId,
		Desc:   r.FormValue("desc"),
	}

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	article := models.Article{
		Title:  request.Title,
		Image:  resp.SecureURL,
		UserID: userId,
		User:   models.UserProfileResponse{},
		Desc:   request.Desc,
	}

	article, err = h.ArticleRepository.CreateArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	article, _ = h.ArticleRepository.GetArticle(article.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: article}
	json.NewEncoder(w).Encode(response)
}

func convertResponseArticle(u models.Article) articledto.ArticleResponse {
	return articledto.ArticleResponse{
		ID:        u.ID,
		Title:     u.Title,
		Image:     u.Image,
		User:      u.User,
		Desc:      u.Desc,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (h *handlerArticle) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(articledto.UpdateArticleRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	article, err := h.ArticleRepository.GetArticle(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		article.Title = request.Title
	}
	if request.Image != "" {
		article.Image = request.Image
	}
	if request.Desc != "" {
		article.Desc = request.Desc
	}

	data, err := h.ArticleRepository.UpdateArticle(article, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticle(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.ArticleRepository.GetArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ArticleRepository.DeleteArticle(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticle(data)}
	json.NewEncoder(w).Encode(response)
}
