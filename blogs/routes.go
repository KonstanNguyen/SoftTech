package blogs

import (
	"SoftTech/blogs/controllers"

	"github.com/gofiber/fiber/v2"
)

// Routes contains all routes relative to /todo
func Routes(app fiber.Router) {
	r := app.Group("/blog")
	r.Get("/list", controllers.GetListBlog)
	r.Post("/update", controllers.UpdateBlog)
	r.Get("/details", controllers.GetDetailBlog)
	r.Post("/create", controllers.CreateBlog)
	r.Post("/delete", controllers.DeleteBlog)
}
