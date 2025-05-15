package models

import (
	"time"

	"github.com/google/uuid"
)

type Route struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	From         uuid.UUID      `gorm:"type:uuid;not null"`
	To           uuid.UUID      `gorm:"type:uuid;not null"`
	EstimatedETA time.Time      `gorm:"not null"`
	Valid        bool           `gorm:"default:true"`

	Segments     []RouteSegment `gorm:"foreignKey:RouteID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RouteSegment struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	RouteID     uuid.UUID `gorm:"type:uuid;not null;index"`
	TransportID uuid.UUID `gorm:"type:uuid;not null"`
	From        uuid.UUID `gorm:"type:uuid;not null"`
	To          uuid.UUID `gorm:"type:uuid;not null"`
	Departure   time.Time `gorm:"not null"`
	Arrival     time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
