package routes

import (
	"milonga/internal/app"
	"milonga/internal/healthcheck"
	"milonga/internal/vigilante"

	"gorm.io/gorm"
)

func RoutesSetup(app *app.App) {

	vigilante.ActivateRoutes(app)
	// alternative to audit users logins
	// vigilante.ActivateRoutes_audit()

	protectedRoutes(app)
	examplesRoutes(app)
	staticRoutes(app)

	HealthRoutes(app)

}

func staticRoutes(app *app.App) {
	app.Server.Static("/", "./web")
}

// get /health and /ping
func HealthRoutes(app *app.App) {

	services := []healthcheck.CriticalService{
		// {
		// 	Name:     "OTher service",
		// 	URL:      "http://other-service:8000/health",
		// 	Timeout:  5 * time.Second,
		// 	Required: true,
		// },
	}

	// checks db connections
	var dbs2Check []*gorm.DB
	app.DB.CheckDefaultConnections()
	dbs2Check = append(dbs2Check, app.DB.Primary)

	healthcheck.ActivateRoutes(app, services, dbs2Check)
}
