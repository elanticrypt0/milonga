package healthcheck

import (
	"milonga/internal/app"

	"gorm.io/gorm"
)

func ActivateRoutes(app *app.App, services []CriticalService, dbs2check []*gorm.DB) {

	healtchecker := NewHealthCheck(app, services, dbs2check)
	app.Server.Get("/health", healtchecker.Check)

}
