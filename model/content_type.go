package model

import "time"

type ContentType struct {
	ID int `gorm: "primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
