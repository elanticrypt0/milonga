package models

import (
	"gorm.io/gorm"
)

// Example model

type File struct {
	ID uint `json:"id" param:"id" query:"id" form:"id"`
	gorm.Model
	Name string `json:"name" param:"name" query:"name" form:"name"`
}

func NewFile() *File {
	return &File{}
}
