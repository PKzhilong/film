package model

type Film struct {
	Name string
	Cover string
	InfoImage string
	Introduce Introduce
	Type int
	Language string
	ShowTime string
	Length int
}

type Introduce struct {
	Director string
	Writer string
	Actors []Actor
	Plot string
}

type Actor struct {
	Name string
}
