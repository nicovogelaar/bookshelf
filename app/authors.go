package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type authorHandler struct {
	repo     bookshelf.AuthorRepository
	validate *validator.Validate
}

func newAuthorHandler(repo bookshelf.AuthorRepository, validate *validator.Validate) authorHandler {
	return authorHandler{repo, validate}
}

func (h *authorHandler) listAuthors(c *gin.Context) {
	p := newPagination(1, 10)

	if err := c.ShouldBindQuery(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authors, err := h.repo.FindAll(p.limit(), p.offset())

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (h *authorHandler) getAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	author, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *authorHandler) createAuthor(c *gin.Context) {
	var author bookshelf.Author

	if err := c.ShouldBind(&author); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&author); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&author); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (h *authorHandler) updateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	author, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := c.ShouldBind(&author); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&author); err != nil {
		log.Debugln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(&author); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *authorHandler) deleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	author, err := h.repo.FindByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := h.repo.Delete(author); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
