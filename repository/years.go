package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type Years struct {
	DB *gorm.DB
}

func (c *Years) CreateIfNotExist(year int) bool {
	var count int
	c.DB.Model(model.Year{}).Where("year = ?", year).Count(&count)
	if count > 0 || year == 0 {
		return false
	}

	newT := &model.Year{
		Year: year,
	}
	c.DB.Create(newT)
	return true
}
