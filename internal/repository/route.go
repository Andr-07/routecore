package repository

import (
	"routecore/internal/models"
	"routecore/pkg/db"

	"github.com/google/uuid"
)

type RouteSegmentDto struct {
	FromID            uuid.UUID 
	ToID              uuid.UUID 
	EarliestDeparture string
	LatestArrival     string
}

type RouteSegmentRepository struct {
	Database *db.Db
}

func NewRouteSegmentRepository(database *db.Db) *RouteSegmentRepository {
	return &RouteSegmentRepository{
		Database: database,
	}
}

func (repo *RouteSegmentRepository) FindActual(query RouteSegmentDto) (*models.RouteSegment, error) {
	var segment models.RouteSegment
	err := repo.Database.DB.
		Where(`"from" = ? AND "to" = ?`, query.FromID, query.ToID).
		Where(`CAST(departure AS time) >= ? AND CAST(arrival AS time) <= ?`, query.EarliestDeparture, query.LatestArrival).
		First(&segment).Error

	if err != nil {
		return nil, err
	}

	return &segment, nil
}

func (repo *RouteSegmentRepository) FindAll() ([]models.RouteSegment, error) {
	var segments []models.RouteSegment
	err := repo.Database.DB.Find(&segments).Error
	if err != nil {
		return nil, err
	}
	return segments, nil
}
