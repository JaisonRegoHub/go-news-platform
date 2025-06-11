package controllers

import (
	"net/http"
	"news-platform/config"
	"news-platform/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	query := config.DB.Preload("Category").Preload("Author").Preload("Tags")

	if categoryID := c.Query("categoryId"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if authorID := c.Query("authorId"); authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}
	if tag := c.Query("tag"); tag != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Joins("JOIN tags ON tags.id = article_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	query.Order("publish_date desc").Find(&articles)
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := config.DB.Preload("Category").Preload("Author").Preload("Tags").First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}
