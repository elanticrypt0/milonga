package handlers

import (
	"fmt"
	"milonga/internal/app"
	"milonga/internal/handlers"
	"milonga/internal/milonga_response"
	"milonga/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminHandler struct {
	handlers.Base
}

// NewUsersHandler crea una nueva instancia de UsersHandler
func NewAdminHandler(app *app.App, DB *gorm.DB) *AdminHandler {
	return &AdminHandler{
		Base: handlers.NewBaseHandler(app, DB),
	}
}

// handler de ejemplo
func (me *AdminHandler) Index(c *fiber.Ctx) error {
	viewPath := fmt.Sprintf("%s/web/admin/index.html", utils.GetAppRootPath())
	return milonga_response.SendHTMLFromFile(c, viewPath)
}
