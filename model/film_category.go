package model

import "time"

type FilmCategory struct {
	ID int `gorm: "primary_key"`

	FilmID int
	CategoryID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
