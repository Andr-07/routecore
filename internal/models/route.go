package models

import (
	"time"

	"github.com/google/uuid"
)

type Route struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	From         uuid.UUID      `gorm:"type:uuid"`
	To           uuid.UUID      `gorm:"type:uuid"`
	EstimatedETA time.Time
	Valid        bool

	Segments []RouteSegment `gorm:"foreignKey:RouteID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RouteSegment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RouteID     uuid.UUID `gorm:"type:uuid;index"`
	TransportID uuid.UUID `gorm:"type:uuid"`
	From        uuid.UUID `gorm:"type:uuid"`
	To          uuid.UUID `gorm:"type:uuid"`
	Departure   time.Time
	Arrival     time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
