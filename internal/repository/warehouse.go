package repository

import (
	"routecore/internal/models"
	"routecore/pkg/db"

	"github.com/google/uuid"
)


type WarehouseRepository struct {
	Database *db.Db
}

func NewWarehouseRepository(database *db.Db) *WarehouseRepository {
	return &WarehouseRepository{
		Database: database,
	}
}

func (repo *WarehouseRepository) FindById(id uuid.UUID) (*models.Warehouse,error) {
	var warehouse models.Warehouse
	result := repo.Database.DB.First(&warehouse, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &warehouse, nil
}