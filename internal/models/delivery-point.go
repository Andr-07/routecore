package models

import (
	"time"

	"github.com/google/uuid"
)

type DeliveryPoint struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Location  GeoPoint  `gorm:"embedded;embeddedPrefix:location_"`
	OpenTime  time.Time
	CloseTime time.Time
	IsActive  bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
