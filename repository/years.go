package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type Years struct {
	DB *gorm.DB
}

func (y *Years) CreateIfNotExist(year int) bool {
	var count int
	y.DB.Model(model.Year{}).Where("year = ?", year).Count(&count)
	if count > 0 || year == 0 {
		return false
	}

	newT := &model.Year{
		Year: year,
	}
	y.DB.Create(newT)
	return true
}

func (y *Years) GetALL() (list []model.Year)  {
	y.DB.Find(&list)
	return
}
