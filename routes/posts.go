package routes

import (
  "github.com/smc181002/jwt-auth/controllers"
  "github.com/smc181002/jwt-auth/middleware"
  "github.com/gofiber/fiber/v2"
)

func PostRoutes(app fiber.Router) {
  api := app.Group("/")
  api.Get("/posts", controllers.GetPost)
  api.Get("/posts/private", middleware.ValidateJWT, controllers.GetPrivatePost)
}
