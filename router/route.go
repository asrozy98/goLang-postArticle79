package router

import (
	"goLang-postArticle79/handler"
	"goLang-postArticle79/repository"
	"goLang-postArticle79/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Word",
		})
	})

	articleRepository := repository.NewArticleRepo(db)
	articleService := service.NewArticleService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService)
	article := route.Group("/article")

	article.POST("/", articleHandler.CreateArticle)
	article.GET("/", articleHandler.GetArticles)
	article.GET("/:id", articleHandler.GetArticle)
	article.PATCH("/:id", articleHandler.UpdateArticle)
	article.DELETE("/:id", articleHandler.DeleteArticle)

	route.Run()
}
