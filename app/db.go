package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // nolint
	"github.com/nicovogelaar/bookshelf/bookshelf"
)

func newDb(config config) *gorm.DB {
	db, err := gorm.Open(config.db.driver, config.db.source)

	if err != nil {
		panic(err)
	}

	db.LogMode(config.db.logMode)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&bookshelf.Author{}, &bookshelf.Book{})

	db.Model(&bookshelf.Book{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")

	return db
}
