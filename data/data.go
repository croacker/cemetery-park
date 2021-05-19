package data

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewSqliteDB - sqlite database
func NewSqliteDB(databaseName string) *DB {

	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	return &DB{db}
}

//Process error
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}