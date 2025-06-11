package utils

import (
	"news-platform/config"
	"news-platform/models"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&models.Category{},
		&models.Author{},
		&models.Tag{},
		&models.Article{},
	)

	SeedData()
}

func SeedData() {
	// Only seed if data is empty
	var count int64
	config.DB.Model(&models.Article{}).Count(&count)
	if count > 0 {
		return
	}

	// Categories
	news := models.Category{Name: "News"}
	tech := models.Category{Name: "Tech"}
	life := models.Category{Name: "Lifestyle"}
	config.DB.Create(&[]models.Category{news, tech, life})

	// Tags
	tagAI := models.Tag{Name: "AI"}
	tagGo := models.Tag{Name: "GoLang"}
	tagVideo := models.Tag{Name: "Video"}
	config.DB.Create(&[]models.Tag{tagAI, tagGo, tagVideo})

	// Authors
	author1 := models.Author{
		Name:        "Alice Johnson",
		ImageURL:    "https://example.com/alice.jpg",
		Description: "Senior Tech Writer",
	}
	author2 := models.Author{
		Name:        "Bob Smith",
		ImageURL:    "https://example.com/bob.jpg",
		Description: "Lifestyle Blogger",
	}
	config.DB.Create(&[]models.Author{author1, author2})

	// Articles
	article1 := models.Article{
		Title:       "Introduction to Go",
		SubTitle:    "Getting started with GoLang",
		ImageURL:    "https://example.com/go-image.jpg",
		ArticleType: "text",
		Description: "<p>Go is an open source programming language...</p>",
		CategoryID:  tech.ID,
		AuthorID:    author1.ID,
		PublishDate: "2025-06-10",
		Tags:        []*models.Tag{&tagGo},
	}

	article2 := models.Article{
		Title:       "AI Trends in 2025",
		SubTitle:    "What's coming next in Artificial Intelligence",
		ImageURL:    "https://example.com/ai.jpg",
		ArticleType: "video",
		MediaURL:    "https://youtube.com/embed/abc123",
		CategoryID:  tech.ID,
		AuthorID:    author1.ID,
		PublishDate: "2025-06-09",
		Tags:        []*models.Tag{&tagAI, &tagVideo},
	}

	article3 := models.Article{
		Title:       "Morning Habits of Successful People",
		SubTitle:    "Improve your lifestyle",
		ImageURL:    "https://example.com/lifestyle.jpg",
		ArticleType: "text",
		Description: "<p>Waking up early can change your life...</p>",
		CategoryID:  life.ID,
		AuthorID:    author2.ID,
		PublishDate: "2025-06-08",
		Tags:        []*models.Tag{},
	}

	config.DB.Create(&[]models.Article{article1, article2, article3})
}
