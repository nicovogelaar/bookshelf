package app

import (
	"github.com/appleboy/gofight"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type authorRepositoryMock struct {
	mock.Mock
}

func (r *authorRepositoryMock) FindByID(id int) (author bookshelf.Author, err error) {
	args := r.Called(id)
	return args.Get(0).(bookshelf.Author), args.Error(1)
}

func (r *authorRepositoryMock) FindAll(limit int, offset int) (authorCollection bookshelf.AuthorCollection, err error) {
	args := r.Called(limit, offset)
	return args.Get(0).(bookshelf.AuthorCollection), args.Error(1)
}

func (r *authorRepositoryMock) Create(author *bookshelf.Author) (err error) {
	args := r.Called(author)
	return args.Error(0)
}

func (r *authorRepositoryMock) Update(author *bookshelf.Author) (err error) {
	args := r.Called(author)
	return args.Error(0)
}

func (r *authorRepositoryMock) Delete(author bookshelf.Author) (err error) {
	args := r.Called(author)
	return args.Error(0)
}

func TestListAuthors(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	authors := bookshelf.AuthorCollection{
		Authors: []bookshelf.Author{
			{
				ID:        1,
				Name:      "Test",
				Biography: "Test 123",
			},
		},
		Limit:      10,
		Offset:     0,
		TotalCount: 1,
	}

	authorRepoMock.On("FindAll", 10, 0).Return(authors, nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.GET("/authors").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"totalCount\":1,\"limit\":10,\"offset\":0,\"authors\":[{\"id\":1,\"name\":\"Test\",\"biography\":\"Test 123\"}]}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	authorRepoMock.AssertExpectations(t)
}

func TestGetAuthor(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	author := bookshelf.Author{
		ID:        1,
		Name:      "Test",
		Biography: "Test 123",
	}

	authorRepoMock.On("FindByID", 1).Return(author, nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.GET("/authors/1").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":1,\"name\":\"Test\",\"biography\":\"Test 123\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	authorRepoMock.AssertExpectations(t)
}

func TestCreateAuthor(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	author := bookshelf.Author{
		Name:      "Test",
		Biography: "Test 123",
	}

	authorRepoMock.On("Create", &author).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.POST("/authors").
		SetDebug(true).
		SetForm(gofight.H{
			"name":      "Test",
			"biography": "Test 123",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":0,\"name\":\"Test\",\"biography\":\"Test 123\"}", r.Body.String())
			assert.Equal(t, http.StatusCreated, r.Code)
		})

	authorRepoMock.AssertExpectations(t)
}

func TestUpdateAuthor(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	author := bookshelf.Author{
		ID:        1,
		Name:      "Test",
		Biography: "Test 123",
	}

	updatedAuthor := bookshelf.Author{
		ID:        1,
		Name:      "Test 123",
		Biography: "Test 456",
	}

	authorRepoMock.On("FindByID", 1).Return(author, nil)
	authorRepoMock.On("Update", &updatedAuthor).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.POST("/authors/1").
		SetDebug(true).
		SetForm(gofight.H{
			"name":      "Test 123",
			"biography": "Test 456",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":1,\"name\":\"Test 123\",\"biography\":\"Test 456\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	authorRepoMock.AssertExpectations(t)
}

func TestDeleteAuthor(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	author := bookshelf.Author{
		ID:        1,
		Name:      "Test",
		Biography: "Test 123",
	}

	authorRepoMock.On("FindByID", 1).Return(author, nil)
	authorRepoMock.On("Delete", author).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.DELETE("/authors/1").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "", r.Body.String())
			assert.Equal(t, http.StatusNoContent, r.Code)
		})

	authorRepoMock.AssertExpectations(t)
}
