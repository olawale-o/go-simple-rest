package src

import (
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth.AuthRoutes(v1)
		protected := v1.Group("/")
		protected.Use(middlewares.Auth())
		articles.ArticleRoutes(protected)
		authors.AuthorRoutes(protected)
	}
}
