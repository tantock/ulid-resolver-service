package server

import (
	"context"

	"github.com/tantock/ulid-resolver-service/internal/dto"

	"github.com/danielgtaylor/huma/v2"
)

func (s *Server) RegisterRoutes(api huma.API) {

	huma.Get(api, "/", s.HelloWorldHandler)

	huma.Get(api, "/health", s.healthHandler)

	huma.Get(api, "/inventory/ulid/upc/{upc}", s.upcToULIDHandler)
}

func (s *Server) HelloWorldHandler(ctx context.Context, input *dto.EmptyInput) (*dto.HelloWorldOutput, error) {
	resp := &dto.HelloWorldOutput{Body: dto.HelloWorldOutputBody{Message: "Hello World"}}
	return resp, nil
}

func (s *Server) healthHandler(ctx context.Context, input *dto.EmptyInput) (*dto.DatabaseHealthOutput, error) {
	health := s.db.Health()

	resp := &dto.DatabaseHealthOutput{
		Body: dto.DatabaseHealthOutputBody{
			Status:            health["status"],
			Message:           health["message"],
			Error:             health["error"],
			OpenConnections:   health["open_connections"],
			InUse:             health["in_use"],
			Idle:              health["idle"],
			WaitCount:         health["wait_count"],
			WaitDuration:      health["wait_duration"],
			MaxIdleClosed:     health["max_idle_closed"],
			MaxLifetimeClosed: health["max_lifetime_closed"],
		},
	}
	return resp, nil
}
