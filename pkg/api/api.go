package api

import (
	"github.com/go-chi/chi"
	"github.com/hscasn/go-microservice/pkg/api/health"
	"github.com/hscasn/go-microservice/pkg/api/ready"
	"github.com/hscasn/go-microservice/pkg/api/settings"
	healthPkg "github.com/hscasn/go-microservice/pkg/health"
)

// Create will bind this API to an existing router
func Create(router *chi.Mux, healthChecks healthPkg.Checks) {
	router.Route("/health", func(r chi.Router) {
		health.Create(r, healthChecks)
	})
	router.Route("/ready", func(r chi.Router) {
		ready.Create(r, healthChecks)
	})
	router.Route("/settings", func(r chi.Router) {
		settings.Create(r)
	})
}
