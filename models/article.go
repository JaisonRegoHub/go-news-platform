package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string   `json:"title"`
	SubTitle    string   `json:"sub_title"`
	ImageURL    string   `json:"image_url"`
	ArticleType string   `json:"article_type"`
	Description string   `json:"description"`
	MediaURL    string   `json:"media_url"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
	AuthorID    uint     `json:"author_id"`
	Author      Author   `json:"author"`
	Tags        []*Tag   `gorm:"many2many:article_tags;" json:"tags"`
	PublishDate string   `json:"publish_date"`
}
