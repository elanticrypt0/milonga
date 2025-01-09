package healthcheck

import (
	"context"
	"fmt"
	"milonga/internal/app"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HealthCheck struct {
	app            *app.App
	Services2Check []CriticalService
	DB2Check       []*gorm.DB
}

func NewHealthCheck(app *app.App, services []CriticalService, dbs []*gorm.DB) *HealthCheck {
	return &HealthCheck{
		app:            app,
		Services2Check: services,
		DB2Check:       dbs,
	}
}

type ServiceStatus struct {
	Healthy bool
	Message string
}

type CriticalService struct {
	Name     string
	URL      string
	Timeout  time.Duration
	Required bool
}

func (me *HealthCheck) CheckDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying sql.DB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %v", err)
	}

	return nil
}

func (me *HealthCheck) checkCriticalServices() ServiceStatus {

	var wg sync.WaitGroup
	results := make(chan error, len(me.Services2Check))

	for _, service := range me.Services2Check {
		wg.Add(1)
		go func(s CriticalService) {
			defer wg.Done()

			client := http.Client{
				Timeout: s.Timeout,
			}

			resp, err := client.Get(s.URL)
			if err != nil {
				if s.Required {
					results <- fmt.Errorf("%s check failed: %v", s.Name, err)
				}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if s.Required {
					results <- fmt.Errorf("%s returned status %d", s.Name, resp.StatusCode)
				}
			}
		}(service)
	}

	wg.Wait()
	close(results)

	var errors []string
	for err := range results {
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return ServiceStatus{
			Healthy: false,
			Message: fmt.Sprintf("Service checks failed: %v", errors),
		}
	}

	return ServiceStatus{
		Healthy: true,
		Message: "All critical services are healthy",
	}
}

func (me *HealthCheck) Check(c *fiber.Ctx) error {
	// Verificar DB
	for _, db := range me.DB2Check {
		if err := me.CheckDBConnection(db); err != nil {
			return c.Status(503).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Database error: %v", err),
			})
		}
	}

	// Verificar servicios críticos
	services := me.checkCriticalServices()
	if !services.Healthy {
		return c.Status(503).JSON(fiber.Map{
			"status":  "error",
			"message": services.Message,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
	})
}
