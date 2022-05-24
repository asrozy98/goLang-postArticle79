package model

import "time"

type Articles struct {
	Id           int       `json:"id" gorm:"primary_key,AUTO_INCREMENT,type:int(11)"`
	Title        string    `json:"title" gorm:"type:varchar(200)"`
	Content      string    `json:"content" gorm:"type:text"`
	Category     string    `json:"category" gorm:"type:varchar(200)"`
	Created_date time.Time `json:"created_date" gorm:"type:datetime default:CURRENT_TIMESTAMP"`
	Updated_date time.Time `json:"updated_date" gorm:"type:datetime default:CURRENT_TIMESTAMP"`
	Status       string    `json:"status" gorm:"type:varchar(100)"`
}

type ArticleRequest struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,enum=Draft0x7CPublish0x7CTrash"`
}
