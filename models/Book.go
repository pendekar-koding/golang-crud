package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
}
