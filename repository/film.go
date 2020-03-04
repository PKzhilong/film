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

func (f Film) UpdateByID(ID int, data *model.Film)   (error) {

	err := f.DB.Model(model.Film{}).Where("id = ?", ID).Update(data).Error
	if err != nil {
		return err
	}
	return nil
}
