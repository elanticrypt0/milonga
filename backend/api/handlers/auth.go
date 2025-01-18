package handlers

import (
	"fmt"
	"milonga/internal/app"
	"milonga/internal/handlers"
	"milonga/internal/milonga_response"
	"milonga/internal/utils"
	"milonga/internal/vigilante"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	handlers.Base
}

// NewUsersHandler crea una nueva instancia de UsersHandler
func NewAuthHandler(app *app.App, DB *gorm.DB) *AuthHandler {
	return &AuthHandler{
		Base: handlers.NewBaseHandler(app, DB),
	}
}

// handler de ejemplo
func (me *AuthHandler) Index(c *fiber.Ctx) error {
	viewPath := fmt.Sprintf("%s/web/auth/index.html", utils.GetAppRootPath())
	return milonga_response.SendHTMLFromFile(c, viewPath)
}

func (me *AuthHandler) Check(c *fiber.Ctx) error {
	// viewPath := fmt.Sprintf("%s/web/auth/index.html", utils.GetAppRootPath())
	// llavar a vigilante y si todo está bien entonces:
	// redirigir hacia /admin/panel
	simpleAuth := vigilante.NewSimpleAuthHandler(me.App, me.App.DB.Primary)
	err := simpleAuth.Login(c)
	if err != nil {
		// return c.SendString(fmt.Sprintf("%s", err))
		return c.Redirect(fmt.Sprintf("%s/admin/auth/", me.App.Config.AppHost))

	}
	return c.Redirect(fmt.Sprintf("%s/admin/panel/", me.App.Config.AppHost))
}

func (me *AuthHandler) Logout(c *fiber.Ctx) error {

	simpleAuth := vigilante.NewSimpleAuthHandler(me.App, me.App.DB.Primary)
	simpleAuth.Logout(c)

	return c.Redirect(fmt.Sprintf("%s/admin/panel/", me.App.Config.AppHost))
}
