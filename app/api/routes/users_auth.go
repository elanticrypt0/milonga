package routes

import (
	"milonga/api/handlers"
	"milonga/internal/app"
	"milonga/internal/middleware"
)

func usersRoutes(app *app.App) {

	middleware := middleware.NewUserAuthMiddelware(app)

	users := app.Server.Group("/users", middleware.IsLogged(), middleware.NotUser())

	handler := handlers.NewUserHandler(app, app.DB.Primary)
	auth_handler := handlers.NewAuthHandler(app, app.DB.Primary)

	users.Get("/", handler.GetAllUsers)

	users.Post("/", handler.CreateUser)

	users.Get("/profile", auth_handler.GetProfile)

	users.Get("/:id", handler.GetUser)
	users.Get("/search", handler.SearchUser)
	users.Put("/:id", handler.UpdateUser)
	users.Delete("/:id", middleware.RequireRole("admin"), handler.DeleteUser)
}

func authRoutes(app *app.App) {
	auth := app.Server.Group("auth")

	handler := handlers.NewAuthHandler(app, app.DB.Primary)

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

}
