package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type HtmlOnline struct {
	DB *gorm.DB
}

func (f HtmlOnline) Create(h *model.HtmlOnline) {
	f.DB.Create(&h)
}

func (f HtmlOnline) Update(h *model.HtmlOnline)  {
	f.DB.Model(model.HtmlOnline{}).Where("id = ?", h.ID).Update(h)
}

func (f HtmlOnline) UpdateByID(id int, h *model.HtmlOnline)  {
	f.DB.Model(model.HtmlOnline{}).Where("id = ?", id).Update(h)
}