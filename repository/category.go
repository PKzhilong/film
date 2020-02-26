package repository

import (
	"filmspider/engine"
	"filmspider/model"
)


type Category struct {

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