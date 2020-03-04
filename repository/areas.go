package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type Areas struct {
	DB *gorm.DB
}

func (c *Areas) CreateIfNotExist(name string) bool {
	var count int
	c.DB.Model(model.Area{}).Where("name = ?", name).Count(&count)
	if count > 0 || len(name) < 1 {
		return false
	}

	newT := &model.Area{
		Name: name,
	}

	c.DB.Create(newT)
	return true
}
