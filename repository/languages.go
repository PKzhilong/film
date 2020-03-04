package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type Languages struct {
	DB *gorm.DB
}

func (c *Languages) CreateIfNotExist(name string) bool {
	var count int
	c.DB.Model(model.Language{}).Where("name = ?", name).Count(&count)
	if count > 0 || len(name) < 1 {
		return false
	}

	newT := &model.Language{
		Name: name,
	}
	c.DB.Create(newT)
	return true
}
