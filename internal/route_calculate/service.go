package route_calculate

import (
	"routecore/internal/models"
	"routecore/internal/repository"

	"github.com/google/uuid"
)

type RouteCalculateService struct {
	DeliveryPointRepository *repository.DeliveryPointRepository
	RouteSegmentRepository  *repository.RouteSegmentRepository
	WarehouseRepository     *repository.WarehouseRepository
}

func NewRouteCalculateService(
	deliveryPointRepository *repository.DeliveryPointRepository,
	routeSegmentRepository *repository.RouteSegmentRepository,
	warehouseRepository *repository.WarehouseRepository,
) *RouteCalculateService {
	return &RouteCalculateService{
		DeliveryPointRepository: deliveryPointRepository,
		RouteSegmentRepository:  routeSegmentRepository,
		WarehouseRepository:     warehouseRepository,
	}
}

func (service *RouteCalculateService) Calculate(from, to uuid.UUID) (*models.RouteSegment, error) {
	warehouse, err := service.WarehouseRepository.FindById(from)
	if err != nil {
		return nil, err
	}

	deliveryPoint, err := service.DeliveryPointRepository.FindById(to)
	if err != nil {
		return nil, err
	}

	dto := repository.RouteSegmentDto{
		FromID:            warehouse.ID,
		ToID:              deliveryPoint.ID,
		EarliestDeparture: warehouse.OpenTime,
		LatestArrival:     deliveryPoint.CloseTime,
	}

	segment, err := service.RouteSegmentRepository.FindActual(dto)
	if err != nil {
		return nil, err
	}

	return segment, nil
}

func (service *RouteCalculateService) GetAll() ([]models.RouteSegment, error) {
	segments, err := service.RouteSegmentRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return segments, nil
}
