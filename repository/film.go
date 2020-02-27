package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type Film struct {
	DB *gorm.DB
}

func (f Film) Create(s *model.Film) {
	f.DB.Create(&s)
}

