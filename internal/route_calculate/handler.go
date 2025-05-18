package route_calculate

import (
	"encoding/json"
	"log"
	"net/http"
	"routecore/configs"
	"routecore/internal/events"
	"routecore/internal/models"
	"routecore/pkg/validation"

	"github.com/google/uuid"
)

type RouteCalculateHandlerDeps struct {
	*configs.Config
	*RouteCalculateService
	*events.EventProducer
}

type RouteCalculateHandler struct {
	*configs.Config
	*RouteCalculateService
	*events.EventProducer
}

func NewRouteCalculateHandler(router *http.ServeMux, deps RouteCalculateHandlerDeps) {
	handler := &RouteCalculateHandler{
		Config:                deps.Config,
		RouteCalculateService: deps.RouteCalculateService,
		EventProducer:         deps.EventProducer,
	}
	router.HandleFunc("POST /routes/calculate", handler.Calculate())
	router.HandleFunc("GET /routes", handler.GetAll())

}

func (handler *RouteCalculateHandler) Calculate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := validation.HandleBody[RouteRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		segment, err := handler.RouteCalculateService.Calculate(body.From, body.To)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		route := models.Route{
			ID:           uuid.New(),
			From:         body.From,
			To:           body.To,
			EstimatedETA: segment.Arrival,
			Valid:        true,
			Segments:     []models.RouteSegment{*segment},
		}

		go func() {
			err := handler.EventProducer.SendRouteCreated(route.ID, route.Segments)
			if err != nil {
				log.Println("❌ Kafka route_created error:", err)
			}
		}()

		resp := RouteResponse{
			ID:           route.ID,
			From:         route.From,
			To:           route.To,
			EstimatedETA: route.EstimatedETA.Format("2006-01-02T15:04:05Z07:00"),
			Valid:        route.Valid,
			Segments:     route.Segments,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (handler *RouteCalculateHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		segments, err := handler.RouteCalculateService.GetAll()
		if err != nil {
			http.Error(w, "failed to fetch routes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(segments); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
