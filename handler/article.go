package handler

import (
	"goLang-postArticle79/model"
	"goLang-postArticle79/service"
	"goLang-postArticle79/utils/validation"
	"strconv"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *articleHandler {
	return &articleHandler{articleService}
}

func (handler *articleHandler) CreateArticle(c *gin.Context) {
	var articleRequest model.ArticleRequest
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		errorMessages := validation.Message(err)

		c.JSON(400, gin.H{
			"success": false,
			"message": "Validation error",
			"error":   errorMessages,
		})
		return
	}

	article, err := handler.articleService.CreateArticle(articleRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Article created",
		"data":    article,
	})
}

func (handler *articleHandler) GetArticles(c *gin.Context) {
	statusString := c.Query("status")
	pageString := c.Query("page")
	page, _ := strconv.Atoi(pageString)
	if page == 0 {
		page = 1
	}

	limitString := c.Query("limit")
	limit, _ := strconv.Atoi(limitString)
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit
	articles, err, total := handler.articleService.GetArticles(offset, limit, statusString)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success":      true,
		"data":         articles,
		"perPageCount": len(articles),
		"allCount":     total,
	})
}

func (handler *articleHandler) UpdateArticle(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var articleRequest model.ArticleRequest
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		errorMessages := validation.Message(err)

		c.JSON(400, gin.H{
			"success": false,
			"message": "Validation error",
			"error":   errorMessages,
		})
		return
	}

	article, err := handler.articleService.UpdateArticle(id, articleRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Article updated",
		"data":    article,
	})
}

func (handler *articleHandler) GetArticle(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	article, err := handler.articleService.GetArticle(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	if article.Id == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Article not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    article,
	})
}

func (handler *articleHandler) DeleteArticle(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := handler.articleService.DeleteArticle(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Article not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Article deleted",
	})
}
