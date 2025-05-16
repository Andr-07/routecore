package repository

import (
	"routecore/internal/dto"
	"routecore/internal/models"
	"routecore/pkg/db"
)

type RouteSegmentRepository struct {
	Database *db.Db
}

func NewRouteSegmentRepository(database *db.Db) *RouteSegmentRepository {
	return &RouteSegmentRepository{
		Database: database,
	}
}

func (repo *RouteSegmentRepository) FindActual(query dto.RouteSegmentDto) (*models.RouteSegment, error) {
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
