package models

import (
	"time"

	"github.com/google/uuid"
)

type DeliveryPoint struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Location  GeoPoint  `gorm:"embedded;embeddedPrefix:location_"`
	OpenTime  string    `gorm:"type:string"`
	CloseTime string    `gorm:"type:string"`
	IsActive  bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
