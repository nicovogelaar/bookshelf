package app

import (
	"github.com/appleboy/gofight"
	"github.com/nicovogelaar/bookshelf/bookshelf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type bookRepositoryMock struct {
	mock.Mock
}

func (r *bookRepositoryMock) FindByID(id int) (book bookshelf.Book, err error) {
	args := r.Called(id)
	return args.Get(0).(bookshelf.Book), args.Error(1)
}

func (r *bookRepositoryMock) FindAll(limit int, offset int) (bookCollection bookshelf.BookCollection, err error) {
	args := r.Called(limit, offset)
	return args.Get(0).(bookshelf.BookCollection), args.Error(1)
}

func (r *bookRepositoryMock) Create(book *bookshelf.Book) (err error) {
	args := r.Called(book)
	return args.Error(0)
}

func (r *bookRepositoryMock) Update(book *bookshelf.Book) (err error) {
	args := r.Called(book)
	return args.Error(0)
}

func (r *bookRepositoryMock) Delete(book bookshelf.Book) (err error) {
	args := r.Called(book)
	return args.Error(0)
}

func TestListBooks(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	books := bookshelf.BookCollection{
		Books: []bookshelf.Book{
			{
				ID:          1,
				Title:       "Test",
				ISBN:        "123456",
				Description: "Test 123",
				AuthorID:    1,
			},
		},
		Limit:      10,
		Offset:     0,
		TotalCount: 1,
	}

	bookRepoMock.On("FindAll", 10, 0).Return(books, nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.GET("/books").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"totalCount\":1,\"limit\":10,\"offset\":0,\"books\":[{\"id\":1,\"authorId\":1,\"author\":{\"id\":0,\"name\":\"\",\"biography\":\"\"},\"title\":\"Test\",\"isbn\":\"123456\",\"description\":\"Test 123\"}]}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	bookRepoMock.AssertExpectations(t)
}

func TestGetBook(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	book := bookshelf.Book{
		ID:          1,
		Title:       "Test",
		Description: "Test 123",
		ISBN:        "123456",
		AuthorID:    1,
	}

	bookRepoMock.On("FindByID", 1).Return(book, nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.GET("/books/1").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":1,\"authorId\":1,\"author\":{\"id\":0,\"name\":\"\",\"biography\":\"\"},\"title\":\"Test\",\"isbn\":\"123456\",\"description\":\"Test 123\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	bookRepoMock.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	book := bookshelf.Book{
		Title:       "Test",
		Description: "Test 123",
		ISBN:        "123456",
		AuthorID:    1,
	}

	bookRepoMock.On("Create", &book).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.POST("/books").
		SetDebug(true).
		SetForm(gofight.H{
			"title":       "Test",
			"description": "Test 123",
			"isbn":        "123456",
			"authorId":    "1",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":0,\"authorId\":1,\"author\":{\"id\":0,\"name\":\"\",\"biography\":\"\"},\"title\":\"Test\",\"isbn\":\"123456\",\"description\":\"Test 123\"}", r.Body.String())
			assert.Equal(t, http.StatusCreated, r.Code)
		})

	bookRepoMock.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	book := bookshelf.Book{
		ID:          1,
		Title:       "Test",
		Description: "Test 123",
		ISBN:        "123456",
		AuthorID:    1,
	}

	updatedBook := bookshelf.Book{
		ID:          1,
		Title:       "Test",
		Description: "Test 456",
		ISBN:        "123456",
		AuthorID:    1,
	}

	bookRepoMock.On("FindByID", 1).Return(book, nil)
	bookRepoMock.On("Update", &updatedBook).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.POST("/books/1").
		SetDebug(true).
		SetForm(gofight.H{
			"title":       "Test",
			"description": "Test 456",
			"isbn":        "123456",
			"authorId":    "1",
		}).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "{\"id\":1,\"authorId\":1,\"author\":{\"id\":0,\"name\":\"\",\"biography\":\"\"},\"title\":\"Test\",\"isbn\":\"123456\",\"description\":\"Test 456\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	bookRepoMock.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	r := gofight.New()

	authorRepoMock := new(authorRepositoryMock)
	bookRepoMock := new(bookRepositoryMock)

	book := bookshelf.Book{
		ID:          1,
		Title:       "Test",
		Description: "Test 123",
	}

	bookRepoMock.On("FindByID", 1).Return(book, nil)
	bookRepoMock.On("Delete", book).Return(nil)

	router := newTestRouter(authorRepoMock, bookRepoMock)

	r.DELETE("/books/1").
		SetDebug(true).
		Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "", r.Body.String())
			assert.Equal(t, http.StatusNoContent, r.Code)
		})

	bookRepoMock.AssertExpectations(t)
}
