package bookshelf

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Author represents an author
type Author struct {
	ID        uint   `gorm:"primary_key" json:"id" form:"id"`
	Name      string `gorm:"index" json:"name" form:"name" validate:"required"`
	Biography string `json:"biography" form:"biography"`
}

// AuthorCollection represents a collection of authors
type AuthorCollection struct {
	TotalCount int      `json:"totalCount"`
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	Authors    []Author `json:"authors"`
}

// AuthorRepository represents an author repository interface
type AuthorRepository interface {
	FindByID(id int) (author Author, err error)
	FindAll(limit int, offset int) (authorCollection AuthorCollection, err error)
	Create(author *Author) (err error)
	Update(author *Author) (err error)
	Delete(author Author) (err error)
}

// DbAuthorRepository represents a database author repository implementation
type DbAuthorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository returns a new instance of author repository
func NewAuthorRepository(db *gorm.DB) *DbAuthorRepository {
	return &DbAuthorRepository{db}
}

// FindByID will find an author by ID
func (r *DbAuthorRepository) FindByID(id int) (author Author, err error) {
	err = r.db.Find(&author, id).Error

	if err != nil {
		log.Errorf("Failed to find author with id '%d': %s", id, err)
	}

	return author, err
}

// FindAll will find all authors
func (r *DbAuthorRepository) FindAll(limit int, offset int) (authorCollection AuthorCollection, err error) {
	var authors []Author
	var totalCount int

	err = r.db.Order("name ASC").Limit(limit).Offset(offset).Find(&authors).Count(&totalCount).Error

	if err != nil {
		log.Errorf("Failed to find all authors: %s", err)
	}

	authorCollection = AuthorCollection{
		TotalCount: totalCount,
		Limit:      limit,
		Offset:     offset,
		Authors:    authors,
	}

	return authorCollection, err
}

// Create will save a new author
func (r *DbAuthorRepository) Create(author *Author) (err error) {
	err = r.db.Create(&author).Error

	if err != nil {
		log.Errorf("Failed to create author: %s", err)
	} else {
		log.WithFields(log.Fields{
			"ID":   author.ID,
			"Name": author.Name,
		}).Debug("Created author")
	}

	return err
}

// Update will save an author
func (r *DbAuthorRepository) Update(author *Author) (err error) {
	err = r.db.Save(&author).Error

	if err != nil {
		log.Errorf("Failed to update author with id '%d': %s", author.ID, err)
	} else {
		log.WithFields(log.Fields{
			"ID":   author.ID,
			"Name": author.Name,
		}).Debug("Updated author")
	}

	return err
}

// Delete will delete an author
func (r *DbAuthorRepository) Delete(author Author) (err error) {
	err = r.db.Delete(author).Error

	if err != nil {
		log.Errorf("Failed to delete author with id '%d': %s", author.ID, err)
	} else {
		log.WithFields(log.Fields{
			"ID":   author.ID,
			"Name": author.Name,
		}).Debug("Deleted author")
	}

	return err
}
