package server

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/joho/godotenv/autoload"

	"github.com/tantock/ulid-resolver-service/internal/database"
)

type Server struct {
	db database.Service
}

// Options for the CLI.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

func NewServer() humacli.CLI {
	NewServer := &Server{
		db: database.New(),
	}
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
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
		r.Mount("/v1", v1)
		config := huma.DefaultConfig("ULID Resolver API", "0.0.0")

		config.Servers = []*huma.Server{
			{
				URL: "/v1",
			},
		}
		api := humachi.New(v1, config)

		NewServer.RegisterRoutes(api)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), r)
		})
	})
	return cli
}
