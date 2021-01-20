package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	Name     string    `json:"name"`
	Artworks []Artwork `json:"artworks"`
}