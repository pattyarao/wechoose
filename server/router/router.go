package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pattyarao/wechoose/controllers"
)

func SetupRoutes(app *fiber.App) {
	//Middleware setting up all the routers to be together
	api := app.Group("/api", logger.New())

	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/signin", controllers.SignIn)

	post := api.Group("/post")
	post.Post("/", controllers.CreatePost)
	post.Get("/:id", controllers.FindPost)
	post.Get("/", controllers.GetPosts)
	post.Get("/user/:userId", controllers.FindUserPosts)
}
