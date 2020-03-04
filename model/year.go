package model

import "time"

type Year struct {
	ID int `gorm: "primary_key"`
	Year int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
