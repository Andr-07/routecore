package repository

import (
	"routecore/internal/models"
	"routecore/pkg/db"

	"github.com/google/uuid"
)


type DeliveryPointRepository struct {
	Database *db.Db
}

func NewDeliveryPointRepository(database *db.Db) *DeliveryPointRepository {
	return &DeliveryPointRepository{
		Database: database,
	}
}

func (repo *DeliveryPointRepository) FindById(id uuid.UUID) (*models.DeliveryPoint,error) {
	var point models.DeliveryPoint
	result := repo.Database.DB.First(&point, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &point, nil
}