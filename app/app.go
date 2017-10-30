package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Run will initialize and start a server
func Run() {
	config := loadConfig()

	if level, err := log.ParseLevel(config.logLevel); err == nil {
		log.SetLevel(level)
	}

	db := newDb(config)
	defer db.Close()

	r := newRouter(config.mode, db)
	r.Run(config.address)
}

func newRouter(mode string, db *gorm.DB) *gin.Engine {
	gin.SetMode(mode)

	r := gin.Default()
	r.Use(newCors())

	ar := bookshelf.NewAuthorRepository(db)
	br := bookshelf.NewBookRepository(db)

	validate := validator.New()

	authorHandler := newAuthorHandler(ar, validate)
	bookHandler := newBookHandler(br, validate)

	addRoutes(r, authorHandler, bookHandler)

	return r
}

func addRoutes(r gin.IRouter, authorHandler authorHandler, bookHandler bookHandler) {
	authors := r.Group("/authors")
	{
		authors.GET("", authorHandler.listAuthors)
		authors.GET("/:id", authorHandler.getAuthor)
		authors.POST("", authorHandler.createAuthor)
		authors.POST("/:id", authorHandler.updateAuthor)
		authors.DELETE("/:id", authorHandler.deleteAuthor)
	}

	books := r.Group("/books")
	{
		books.GET("", bookHandler.listBooks)
		books.GET("/:id", bookHandler.getBook)
		books.POST("", bookHandler.createBook)
		books.POST("/:id", bookHandler.updateBook)
		books.DELETE("/:id", bookHandler.deleteBook)
	}
}

func newCors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, http.MethodDelete)
	corsConfig.AllowAllOrigins = true

	return cors.New(corsConfig)
}
