package routes

import (
	"news-platform/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/categories", controllers.GetCategories)
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/:id", controllers.GetArticle)
	r.GET("/authors/:id", controllers.GetAuthor)

	return r
}
