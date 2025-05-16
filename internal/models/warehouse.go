package models

import (
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Location  GeoPoint `gorm:"embedded;embeddedPrefix:location_"`
	OpenTime  string   `gorm:"type:string"`
	CloseTime string   `gorm:"type:string"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
