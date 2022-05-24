package main

import (
	"goLang-postArticle79/config"
	"goLang-postArticle79/model"
	"goLang-postArticle79/router"
	"goLang-postArticle79/utils"
)

func main() {
	utils.Utils()
	db := config.DatabaseConfig()
	db.AutoMigrate(&model.Articles{})

	router.Router(db)
}
