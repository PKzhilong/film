package repository

import (
	"filmspider/engine"
	"filmspider/model"
	"github.com/jinzhu/gorm"
)


type Category struct {
	DB *gorm.DB
}

func (c Category) GetCategoryByName(name string) (res model.Category) {
	var db = engine.DBRun()
	db.Where("name like ?", "%" + name + "%").First(&res)
	return
}

func (c Category) CreateCategory(content *model.Category)  {
	var db = engine.DBRun()
	db.Create(content)
}

func (c Category) CreateIfNotExist(name string) bool {
	var res model.Category
	c.DB.Where("name like ?", "%" + name + "%").First(&res)
	if res.ID != 0 {
		return false
	}

	content := &model.Category{
		Name: name,
	}
	c.DB.Create(content)
	return true
}
