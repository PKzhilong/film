package repository

import (
	"container/list"
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type ContentTypes struct {
	DB *gorm.DB
}

func (c *ContentTypes) CreateIfExist(name string) int {
	var contentType model.ContentType
	c.DB.Model(model.ContentType{}).Where("name like ?", "%" + name + "%").First(&contentType)
	if contentType.ID > 0 {
		return contentType.ID
	}

	newT := &model.ContentType{
		Name: name,
	}
	c.DB.Create(newT)
	return newT.ID
}

func (c *ContentTypes) GetAll() (list.Element) {

}
