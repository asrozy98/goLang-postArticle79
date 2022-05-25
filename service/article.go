package service

import (
	"goLang-postArticle79/model"
	"goLang-postArticle79/repository"
	"time"
)

type ArticleService interface {
	CreateArticle(articleRequest model.ArticleRequest) (model.Articles, error)
	GetArticles(offset int, limit int, status string) ([]model.Articles, error, int64)
	GetArticle(id int) (model.Articles, error)
	UpdateArticle(id int, articleRequest model.ArticleRequest) (model.Articles, error)
	DeleteArticle(id int) error
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) *articleService {
	return &articleService{repository}
}

func (s *articleService) CreateArticle(articleRequest model.ArticleRequest) (model.Articles, error) {
	article := model.Articles{
		Title:        articleRequest.Title,
		Content:      articleRequest.Content,
		Category:     articleRequest.Category,
		Status:       articleRequest.Status,
		Created_date: time.Now(),
		Updated_date: time.Now(),
	}
	newArticle, err := s.articleRepository.CreateArticle(article)
	return newArticle, err
}

func (s *articleService) GetArticles(offset int, limit int, status string) ([]model.Articles, error, int64) {
	return s.articleRepository.GetArticles(offset, limit, status)
}

func (s *articleService) GetArticle(id int) (model.Articles, error) {
	return s.articleRepository.GetArticle(id)
}

func (s *articleService) UpdateArticle(id int, articleRequest model.ArticleRequest) (model.Articles, error) {
	article, err := s.articleRepository.GetArticle(id)
	if err != nil {
		return article, err
	}

	article.Title = articleRequest.Title
	article.Content = articleRequest.Content
	article.Category = articleRequest.Category
	article.Status = articleRequest.Status
	article.Updated_date = time.Now()

	newArticle, err := s.articleRepository.UpdateArticle(article)
	return newArticle, err
}

func (s *articleService) DeleteArticle(id int) error {
	return s.articleRepository.DeleteArticle(id)
}
