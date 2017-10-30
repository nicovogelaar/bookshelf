package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type bookHandler struct {
	repo     bookshelf.BookRepository
	validate *validator.Validate
}

func newBookHandler(repo bookshelf.BookRepository, validate *validator.Validate) bookHandler {
	return bookHandler{repo, validate}
}

func (h *bookHandler) listBooks(c *gin.Context) {
	p := newPagination(1, 10)

	if err := c.ShouldBindQuery(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	books, err := h.repo.FindAll(p.limit(), p.offset())

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *bookHandler) getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	book, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) createBook(c *gin.Context) {
	var book bookshelf.Book

	if err := c.ShouldBind(&book); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&book); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&book); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *bookHandler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	book, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := c.ShouldBind(&book); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&book); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(&book); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	book, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := h.repo.Delete(book); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
