package dto

import (
	"time"

	"github.com/google/uuid"
)

type RouteSegmentDto struct {
	FromID            uuid.UUID 
	ToID              uuid.UUID 
	EarliestDeparture time.Time 
	LatestArrival     time.Time 
}