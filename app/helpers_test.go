package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	"gopkg.in/go-playground/validator.v9"
)

func newTestRouter(ar bookshelf.AuthorRepository, br bookshelf.BookRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(newCors())

	validate := validator.New()

	authorHandler := newAuthorHandler(ar, validate)
	bookHandler := newBookHandler(br, validate)

	addRoutes(r, authorHandler, bookHandler)

	return r
}
