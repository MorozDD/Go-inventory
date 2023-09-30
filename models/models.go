package models

import (
	"github.com/jinzhu/gorm"
)

type Part struct {
	gorm.Model
	Name              string            `json:"name"`
	SKU               string            `json:"sku"`
	Description       string            `json:"description"`
	Price             float64           `json:"price"`
	Location          string            `json:"location"`
	ShipmentPackaging ShipmentPackaging `json:"shipment_packaging" gorm:"type:jsonb"`
	Attributes        []Attribute       `json:"attributes" gorm:"many2many:part_attributes"`
	Fitments          []Fitment         `json:"fitments"`
	Images            []Image           `json:"images"`
	Metadata          []Metadata        `json:"metadata"`
	Version           int               `json:"version"`
}

type Attribute struct {
	gorm.Model
	Name  string `json:"name"`
	Value string `json:"value"`
	Parts []Part `gorm:"many2many:part_attributes"`
}

type Fitment struct {
	gorm.Model
	PartID   uint   `json:"-"`
	Year     int    `json:"year"`
	Make     string `json:"make"`
	CarModel string `json:"model"`
}

type Image struct {
	gorm.Model
	PartID   uint   `json:"-"`
	ImageURL string `json:"image_url"`
}

type Metadata struct {
	gorm.Model
	PartID uint   `json:"-"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type Location struct {
	gorm.Model
	Name  string `json:"name"`
	Parts []Part `gorm:"many2many:part_locations"`
}

type ShipmentPackaging struct {
	Weight    float64 `json:"weight"`
	Size      string  `json:"size"`
	Hazardous bool    `json:"hazardous"`
	Fragile   bool    `json:"fragile"`
}
