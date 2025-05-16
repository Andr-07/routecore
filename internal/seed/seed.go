package seed

import (
	"routecore/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	wh1 := models.Warehouse{
		ID:   uuid.New(),
		Name: "Склад Москва",
		Location: models.GeoPoint{
			Latitude:  55.751244,
			Longitude: 37.618423,
		},
		OpenTime:  "07:00:00",
		CloseTime: "22:00:00",
	}

	wh2 := models.Warehouse{
		ID:   uuid.New(),
		Name: "Склад Санкт-Петербург",
		Location: models.GeoPoint{
			Latitude:  59.934280,
			Longitude: 30.335098,
		},
		OpenTime:  "08:00:00",
		CloseTime: "21:00:00",
	}

	dp1 := models.DeliveryPoint{
		ID: uuid.New(),
		Location: models.GeoPoint{
			Latitude:  55.7601,
			Longitude: 37.6185,
		},
		OpenTime:  "08:00:00",
		CloseTime: "20:00:00",
		IsActive:  true,
	}

	dp2 := models.DeliveryPoint{
		ID: uuid.New(),
		Location: models.GeoPoint{
			Latitude:  59.9386,
			Longitude: 30.3141,
		},
		OpenTime:  "08:00:00",
		CloseTime: "21:00:00",
		IsActive:  true,
	}

	dp3 := models.DeliveryPoint{
		ID: uuid.New(),
		Location: models.GeoPoint{
			Latitude:  56.837,
			Longitude: 60.597,
		},
		OpenTime:  "10:00:00",
		CloseTime: "20:00:00",
		IsActive:  false,
	}

	rs1 := models.RouteSegment{
		TransportID: uuid.New(),
		From:        wh1.ID,
		To:          dp1.ID,
		Departure:   time.Now().Add(time.Hour * 2),
		Arrival:     time.Now().Add(time.Hour * 5),
	}

	route := models.Route{
		ID:           uuid.New(),
		From:         wh1.ID,
		To:           dp1.ID,
		Segments:     []models.RouteSegment{rs1},
		EstimatedETA: rs1.Arrival,
		Valid:        true,
	}

	if err := db.Create(&wh1).Error; err != nil {
		return err
	}
	if err := db.Create(&wh2).Error; err != nil {
		return err
	}
	if err := db.Create(&dp1).Error; err != nil {
		return err
	}
	if err := db.Create(&dp2).Error; err != nil {
		return err
	}
	if err := db.Create(&dp3).Error; err != nil {
		return err
	}
	if err := db.Create(&route).Error; err != nil {
		return err
	}
	for _, seg := range route.Segments {
		seg.RouteID = route.ID
		if err := db.Create(&seg).Error; err != nil {
			return err
		}
	}

	return nil
}
