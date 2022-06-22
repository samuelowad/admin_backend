package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/controllers"
	"github.com/samuelowad/admin_backend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)

	app.Use(middleware.IsAuth)
	app.Post("/api/v1/logout", controllers.Logout)

	//
	app.Get("/api/v1/user", controllers.User)
	app.Get("/api/v1/users", controllers.AllUsers)
	app.Post("/api/v1/create-user", controllers.CreateUser)
	app.Get("/api/v1/users/:id", controllers.GetUser)
	app.Put("/api/v1/users/:id", controllers.UpdateUser)
	app.Delete("/api/v1/users/:id", controllers.DeleteUser)

	app.Get("/api/v1/roles", controllers.AllRole)
	app.Post("/api/v1/create-role", controllers.CreateRole)
	app.Get("/api/v1/roles/:id", controllers.GetRole)
	app.Put("/api/v1/roles/:id", controllers.UpdateRole)
	app.Delete("/api/v1/roles/:id", controllers.DeleteRole)

	app.Get("/api/v1/permissions", controllers.AllPermissions)

	app.Get("/api/v1/products", controllers.AllProducts)
	app.Post("/api/v1/create-product", controllers.CreateProduct)
	app.Get("/api/v1/products/:id", controllers.GetProduct)
	app.Put("/api/v1/products/:id", controllers.UpdateProduct)
	app.Delete("/api/v1/products/:id", controllers.DeleteProduct)

	app.Post("/api/v1/upload", controllers.ImageUpload)

	app.Static("/api/uploads", "./uploads")

}
