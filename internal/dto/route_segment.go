package dto

import (
	"github.com/google/uuid"
)

type RouteSegmentDto struct {
	FromID            uuid.UUID 
	ToID              uuid.UUID 
	EarliestDeparture string
	LatestArrival     string
}