package events

import (
	"encoding/json"
	"routecore/internal/models"
	"routecore/pkg/kafka"
	"time"

	"github.com/google/uuid"
)

type EventProducer struct {
	KafkaProducer *kafka.KafkaProducer
}

func NewEventProducer(kafkaProducer *kafka.KafkaProducer) *EventProducer {
	return &EventProducer{
		KafkaProducer: kafkaProducer,
	}
}

func (e *EventProducer) SendRouteCreated(routeID uuid.UUID, segments []models.RouteSegment) error {
	event := map[string]interface{}{
		"event_type": "route_created",
		"route_id":   routeID,
		"segments":   segments,
		"timestamp":  time.Now().UTC(),
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return e.KafkaProducer.WriteMessage(data)
}

func (e *EventProducer) SendRouteUpdated(routeID uuid.UUID, reason string) error {
	event := map[string]interface{}{
		"event_type": "route_updated",
		"route_id":   routeID,
		"reason":     reason,
		"timestamp":  time.Now().UTC(),
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return e.KafkaProducer.WriteMessage(data)
}