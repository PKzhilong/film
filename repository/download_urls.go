package repository

import (
	"filmspider/model"
	"github.com/jinzhu/gorm"
)

type DownloadUrls struct {
	DB *gorm.DB
}

func (f DownloadUrls) Create(h *model.DownloadUrl) {
	f.DB.Create(&h)
}

func (f DownloadUrls) Update(h *model.DownloadUrl)  {
	f.DB.Model(model.DownloadUrl{}).Where("id = ?", h.ID).Update(h)
}
