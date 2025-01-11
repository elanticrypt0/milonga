package vigilante

import (
	"milonga/internal/app"
)

func ActivateRoutes(app *app.App) {

	usersRoutes(app)
	authRoutes(app)

}

func usersRoutes(app *app.App) {

	middleware := NewVigilanteMiddelware(app)

	users := app.Server.Group("/users", middleware.IsLogged(), middleware.NotUser())

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
	users.Post("/new/vip/:email", middleware.RequireRole("admin"), handler.CreateVIPGuest)

}

func authRoutes(app *app.App) {
	auth := app.Server.Group("auth")

	handler := NewAuthHandler(app, app.DB.Primary)

	auth.Post("/register", handler.Register)
	auth.Post("/login/guest", handler.LoginByPasswordToken)
	auth.Post("/login", handler.Login)

}