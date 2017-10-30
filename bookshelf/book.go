package bookshelf

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Book represents a book
type Book struct {
	ID          uint   `gorm:"primary_key" json:"id" form:"id"`
	AuthorID    uint   `json:"authorId" form:"authorId" validate:"required"`
	Author      Author `json:"author" form:"author" validate:"structonly"`
	Title       string `gorm:"index" json:"title" form:"title" validate:"required"`
	ISBN        string `json:"isbn" form:"isbn"`
	Description string `json:"description" form:"description"`
}

// BookCollection represents a collection of books
type BookCollection struct {
	TotalCount int    `json:"totalCount"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Books      []Book `json:"books"`
}

// BookRepository represents a book repository interface
type BookRepository interface {
	FindByID(id int) (book Book, err error)
	FindAll(limit int, offset int) (bookCollection BookCollection, err error)
	Create(book *Book) (err error)
	Update(book *Book) (err error)
	Delete(book Book) (err error)
}

// DbBookRepository represents a database book repository implementation
type DbBookRepository struct {
	db *gorm.DB
}

// NewBookRepository returns a new instance of book repository
func NewBookRepository(db *gorm.DB) *DbBookRepository {
	return &DbBookRepository{db}
}

// FindByID will find a book by ID
func (r *DbBookRepository) FindByID(id int) (book Book, err error) {
	err = r.db.Find(&book, id).Error

	if err != nil {
		log.Errorf("Failed to find book with id '%d': %s", id, err)
	}

	return book, err
}

// FindAll will find all books
func (r *DbBookRepository) FindAll(limit int, offset int) (bookCollection BookCollection, err error) {
	var books []Book
	var totalCount int

	err = r.db.Preload("Author").Order("title ASC").Limit(limit).Offset(offset).Find(&books).Count(&totalCount).Error

	if err != nil {
		log.Errorf("Failed to find all books: %s", err)
	}

	bookCollection = BookCollection{
		TotalCount: totalCount,
		Limit:      limit,
		Offset:     offset,
		Books:      books,
	}

	return bookCollection, err
}

// Create will save a new book
func (r *DbBookRepository) Create(book *Book) (err error) {
	err = r.db.Create(&book).Error

	if err != nil {
		log.Errorf("Failed to create book: %s", err)
	} else {
		log.WithFields(log.Fields{
			"ID":       book.ID,
			"Title":    book.Title,
			"AuthorID": book.AuthorID,
		}).Debug("Created book")
	}

	return err
}

// Update will save a book
func (r *DbBookRepository) Update(book *Book) (err error) {
	err = r.db.Save(&book).Error

	if err != nil {
		log.Errorf("Failed to update book with id '%d': %s", book.ID, err)
	} else {
		log.WithFields(log.Fields{
			"ID":       book.ID,
			"Title":    book.Title,
			"AuthorID": book.AuthorID,
		}).Debug("Updated book")
	}

	return err
}

// Delete will delete a book
func (r *DbBookRepository) Delete(book Book) (err error) {
	err = r.db.Delete(book).Error

	if err != nil {
		log.Errorf("Failed to delete book with id '%d': %s", book.ID, err)
	} else {
		log.WithFields(log.Fields{
			"ID":    book.ID,
			"Title": book.Title,
		}).Debug("Deleted book")
	}

	return err
}
