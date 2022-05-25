package repository

import (
	"goLang-postArticle79/model"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article model.Articles) (model.Articles, error)
	GetArticles(offset int, limit int, status string) ([]model.Articles, error, int64)
	GetArticle(ID int) (model.Articles, error)
	UpdateArticle(article model.Articles) (model.Articles, error)
	DeleteArticle(ID int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticles(offset int, limit int, status string) ([]model.Articles, error, int64) {
	var articles []model.Articles
	var count int64 = 0
	if status != "" {
		err := r.db.Where("status = ?", status).Offset(offset).Limit(limit).Find(&articles).Count(&count).Error
		return articles, err, count
	}
	r.db.Model(&articles).Count(&count)
	err := r.db.Limit(limit).Offset(offset).Find(&articles).Error
	return articles, err, count
}

func (r *articleRepository) CreateArticle(article model.Articles) (model.Articles, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *articleRepository) GetArticle(ID int) (model.Articles, error) {
	var article model.Articles
	err := r.db.Find(&article, ID).Error
	return article, err
}

func (r *articleRepository) UpdateArticle(article model.Articles) (model.Articles, error) {
	err := r.db.Save(&article).Error
	return article, err
}

func (r *articleRepository) DeleteArticle(ID int) error {
	err := r.db.Delete(&model.Articles{}, ID).Error
	return err
}
