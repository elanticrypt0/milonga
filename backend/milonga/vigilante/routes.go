package vigilante

import (
	"milonga/milonga/app"

	"github.com/gofiber/fiber/v2"
)

func ActivateRoutes(app *app.App, router fiber.Router) {

	usersRoutes(app, router)
	authRoutes(app, router)

}

func usersRoutes(app *app.App, router fiber.Router) {

	middleware := NewVigilanteMiddelware(app)

	users := router.Group("/users", middleware.IsLogged(), middleware.IsStaff())

	handler := NewUserHandler(app, app.DB.Primary)
	auth_handler := NewAuthHandler(app, app.DB.Primary)

	users.Get("/", handler.GetAllUsers)

	users.Post("/", handler.CreateUser)

	users.Get("/profile", auth_handler.GetProfile)

	users.Get("/:id", handler.GetUser)
	users.Get("/search", handler.SearchUser)
	users.Put("/:id", handler.UpdateUser)
	users.Delete("/:id", middleware.RequireRole("admin"), handler.DeleteUser)

	// add VIPGUEST
	// User with password token
	users.Post("/new/guest", middleware.RequireRole("admin"), handler.CreateAccess2TokenLogin)

}

func authRoutes(app *app.App, router fiber.Router) {
	auth := router.Group("auth")

	handler := NewAuthHandler(app, app.DB.Primary)

	auth.Post("/register", handler.Register)
	auth.Get("/login/guest/link", handler.LoginByPasswordTokenWithLink)
	auth.Post("/login/guest", handler.LoginByPasswordToken)
	auth.Post("/login", handler.Login)

}
