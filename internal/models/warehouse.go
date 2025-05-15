package models

import (
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Location  GeoPoint  `gorm:"embedded;embeddedPrefix:location_"`
	OpenTime  time.Time
	CloseTime time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}