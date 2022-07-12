package routes

import (
  "github.com/gofiber/fiber/v2"
  "github.com/smc181002/jwt-auth/controllers"
)

func AuthRoutes(app fiber.Router) {
  api := app.Group("/auth")
  api.Post("/", controllers.AuthenticateUser)
}
