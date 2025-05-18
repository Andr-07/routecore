package main

import (
	"log"
	"net/http"
	"routecore/configs"
	"routecore/internal/events"
	"routecore/internal/repository"
	"routecore/internal/route_calculate"
	"routecore/pkg/db"
	"routecore/pkg/kafka"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(&conf.Db)

	kafkaProducer := kafka.NewKafkaProducer(&conf.Kafka)
	defer kafkaProducer.Writer.Close()

	router := http.NewServeMux()

	// Events
	eventProducer := events.NewEventProducer(kafkaProducer)

	// Repositories
	deliveryPointRepository := repository.NewDeliveryPointRepository(db)
	routeSegmentRepository := repository.NewRouteSegmentRepository(db)
	warehouseRepository := repository.NewWarehouseRepository(db)

	// Services
	routeCalculateService := route_calculate.NewRouteCalculateService(
		deliveryPointRepository,
		routeSegmentRepository,
		warehouseRepository,
	)

	// Handlers
	route_calculate.NewRouteCalculateHandler(router, route_calculate.RouteCalculateHandlerDeps{
		Config:                conf,
		RouteCalculateService: routeCalculateService,
		EventProducer: eventProducer,
	})

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
