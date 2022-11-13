package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticles() ([]models.Article, error)
	GetArticle(ID int) (models.Article, error)
	CreateArticle(Article models.Article) (models.Article, error)
	UpdateArticle(Article models.Article, ID int) (models.Article, error)
	DeleteArticle(Article models.Article, ID int) (models.Article, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Preload("User").Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.Preload("User").First(&article, ID).Error

	return article, err
}

func (r *repository) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}

func (r *repository) UpdateArticle(article models.Article, ID int) (models.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}
func (r *repository) DeleteArticle(article models.Article, ID int) (models.Article, error) {
	err := r.db.Delete(&article).Error

	return article, err
}
