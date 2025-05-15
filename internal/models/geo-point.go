package models

type GeoPoint struct {
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
}