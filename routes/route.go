package routes

import (
	"react_go_catalog_web/controllers/imagecontrollers"
	"react_go_catalog_web/controllers/postcontrollers"
	"react_go_catalog_web/controllers/usercontrollers"
	"react_go_catalog_web/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", usercontrollers.Register)
	app.Post("/api/login", usercontrollers.Login)

	app.Use(middleware.IsAuthenticate)

	app.Post("/api/content", postcontrollers.CreatePost)
	app.Get("/api/content", postcontrollers.ReadPost)
	app.Get("/api/content/:id", postcontrollers.DetailPost)
	app.Put("/api/update/update/:id", postcontrollers.UpdatePost)
	app.Get("/api/unique", postcontrollers.UniquePost)
	app.Delete("/api/content/delete/:id", postcontrollers.DeletePost)

	app.Post("/api/content/upload-image", imagecontrollers.ImageUpload)
	app.Static("/api/content/uploads", "./uploads")
}
