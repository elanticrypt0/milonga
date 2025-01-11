package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateModel(modelName string) error {
	// Convert model name to proper case
	modelName = strings.Title(strings.ToLower(modelName))

	// Generate model file
	if err := generateModelFile(modelName); err != nil {
		return fmt.Errorf("error generating model: %v", err)
	}

	// Generate handler file
	if err := generateHandlerFile(modelName); err != nil {
		return fmt.Errorf("error generating handler: %v", err)
	}

	// Generate routes file
	if err := generateRoutesFile(modelName); err != nil {
		return fmt.Errorf("error generating routes: %v", err)
	}

	fmt.Printf("Successfully generated %s model with CRUD operations\n", modelName)
	return nil
}

func generateModelFile(modelName string) error {
	modelTemplate := `package models

import (
	"gorm.io/gorm"
)

type {{.Name}} struct {
	gorm.Model
	// Add your fields here
}
`
	return generateFile(
		filepath.Join("api", "models", strings.ToLower(modelName)+".go"),
		modelTemplate,
		struct{ Name string }{Name: modelName},
	)
}

func generateHandlerFile(modelName string) error {
	handlerTemplate := `package handlers

import (
	"milonga/api/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type {{.Name}}Handler struct {
	db *gorm.DB
}

func New{{.Name}}Handler(db *gorm.DB) *{{.Name}}Handler {
	return &{{.Name}}Handler{db: db}
}

func (h *{{.Name}}Handler) Create(c *fiber.Ctx) error {
	var item models.{{.Name}}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.db.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *{{.Name}}Handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.{{.Name}}

	if err := h.db.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}

	return c.JSON(item)
}

func (h *{{.Name}}Handler) List(c *fiber.Ctx) error {
	var items []models.{{.Name}}
	
	if err := h.db.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(items)
}

func (h *{{.Name}}Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.{{.Name}}

	if err := h.db.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}

	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.db.Save(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(item)
}

func (h *{{.Name}}Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.{{.Name}}

	if err := h.db.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}

	if err := h.db.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Record deleted successfully",
	})
}
`
	return generateFile(
		filepath.Join("api", "handlers", strings.ToLower(modelName)+"_handler.go"),
		handlerTemplate,
		struct{ Name string }{Name: modelName},
	)
}

func generateRoutesFile(modelName string) error {
	routesTemplate := `package routes

import (
	"milonga/api/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup{{.Name}}Routes(app *fiber.App, db *gorm.DB) {
	handler := handlers.New{{.Name}}Handler(db)

	routes := app.Group("/{{.LowerName}}")
	
	routes.Post("/", handler.Create)
	routes.Get("/:id", handler.Get)
	routes.Get("/", handler.List)
	routes.Put("/:id", handler.Update)
	routes.Delete("/:id", handler.Delete)
}
`
	return generateFile(
		filepath.Join("api", "routes", strings.ToLower(modelName)+"_routes.go"),
		routesTemplate,
		struct {
			Name      string
			LowerName string
		}{
			Name:      modelName,
			LowerName: strings.ToLower(modelName),
		},
	)
}

func generateFile(path string, tmpl string, data interface{}) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Parse and execute template
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	if err := t.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}
