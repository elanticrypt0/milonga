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

	// update otp maxuses
	users.Put("/otp/", middleware.RequireRole("admin"), handler.UpdateOTPMaxUses)
	// this uses can use just one time password
	users.Post("/otp/new/guest", middleware.RequireRole("admin"), handler.CreateAccess2TokenLogin)

	users.Post("/otp/new/vip", middleware.RequireRole("admin"), handler.CreateAccess2TokenLogin)

}

func authRoutes(app *app.App, router fiber.Router) {
	auth := router.Group("auth")

	handler := NewAuthHandler(app, app.DB.Primary)

	auth.Post("/register", handler.Register)
	auth.Get("/login/otp/link", handler.LoginByPasswordTokenWithLink)
	auth.Post("/login/otp", handler.LoginByPasswordToken)
	auth.Post("/login", handler.Login)

}
