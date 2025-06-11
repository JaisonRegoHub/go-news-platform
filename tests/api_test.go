package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"news-platform/config"
	"news-platform/models"
	"news-platform/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestRouter() http.Handler {
	config.ConnectDB()
	config.DB.Exec("DROP TABLE IF EXISTS article_tags;")
	config.DB.Migrator().DropTable(&models.Article{}, &models.Category{}, &models.Tag{}, &models.Author{})
	config.DB.AutoMigrate(&models.Category{}, &models.Tag{}, &models.Author{}, &models.Article{})

	// Seed minimal test data
	category := models.Category{Name: "Test Category"}
	tag := models.Tag{Name: "Test Tag"}
	author := models.Author{Name: "Test Author", ImageURL: "https://image.com", Description: "Desc"}
	config.DB.Create(&category)
	config.DB.Create(&tag)
	config.DB.Create(&author)

	article := models.Article{
		Title:       "Test Article",
		SubTitle:    "Sub",
		ImageURL:    "http://image.com",
		ArticleType: "text",
		Description: "<p>Some HTML</p>",
		CategoryID:  category.ID,
		AuthorID:    author.ID,
		PublishDate: "2025-06-11",
		Tags:        []*models.Tag{&tag},
	}
	config.DB.Create(&article)

	return routes.SetupRoutes()
}

func TestGetCategories(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/categories", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	var categories []models.Category
	err := json.Unmarshal(resp.Body.Bytes(), &categories)
	assert.Nil(t, err)
	assert.Greater(t, len(categories), 0)
}

func TestGetArticles(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/articles", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	var articles []models.Article
	err := json.Unmarshal(resp.Body.Bytes(), &articles)
	assert.Nil(t, err)
	assert.Greater(t, len(articles), 0)
}

func TestGetArticleByID(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	var article models.Article
	err := json.Unmarshal(resp.Body.Bytes(), &article)
	assert.Nil(t, err)
	assert.Equal(t, "Test Article", article.Title)
}

func TestGetAuthorByID(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/authors/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	var author models.Author
	err := json.Unmarshal(resp.Body.Bytes(), &author)
	assert.Nil(t, err)
	assert.Equal(t, "Test Author", author.Name)
}
