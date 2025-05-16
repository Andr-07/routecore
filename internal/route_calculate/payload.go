package route_calculate

import (
	"routecore/internal/models"

	"github.com/google/uuid"
)

type RouteRequest struct {
	From uuid.UUID `json:"from" validate:"required,uuid"`
	To   uuid.UUID `json:"to" validate:"required,uuid"`
}

type RouteResponse struct {
	ID           uuid.UUID             `json:"id"`
	From         uuid.UUID             `json:"from"`
	To           uuid.UUID             `json:"to"`
	EstimatedETA string                `json:"estimated_eta"`
	Valid        bool                  `json:"valid"`
	Segments     []models.RouteSegment `json:"segments"`
}
