package route_calculate

import (
	"encoding/json"
	"net/http"
	"routecore/configs"
	"routecore/internal/models"
	"routecore/pkg/validation"

	"github.com/google/uuid"
)

type RouteCalculateHandlerDeps struct {
	*configs.Config
	*RouteCalculateService
}

type RouteCalculateHandler struct {
	*configs.Config
	*RouteCalculateService
}

func NewRouteCalculateHandler(router *http.ServeMux, deps RouteCalculateHandlerDeps) {
	handler := &RouteCalculateHandler{
		Config:                deps.Config,
		RouteCalculateService: deps.RouteCalculateService,
	}
	router.HandleFunc("POST /route/calculate", handler.Calculate())

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

		resp := RouteResponse{
			ID:           route.ID,
			From:         route.From,
			To:           route.To,
			EstimatedETA: route.EstimatedETA.Format("2006-01-02T15:04:05Z07:00"),
			Valid:        route.Valid,
			Segments:     route.Segments,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
