package models

import (
	"github.com/LeonardoRafaelli/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func createBook(b *Book) *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
