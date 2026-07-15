package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/tantock/ulid-resolver-service/internal/dto"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()
	config := huma.DefaultConfig("ULID Resolver API", "0.0.0")

	config.Servers = []*huma.Server{
		{
			URL: "/v1",
		},
	}
	api := humachi.New(v1, config)

	huma.Get(api, "/", s.HelloWorldHandler)

	huma.Get(api, "/health", s.healthHandler)

	r.Mount("/v1", v1)
	return r
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
