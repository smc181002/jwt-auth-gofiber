package main

import (
  "log"
  "github.com/gofiber/fiber/v2"
  "github.com/smc181002/jwt-auth/routes"
  "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
  app := fiber.New()
  app.Use(recover.New())

  api := app.Group("/api")

  // Post Routes
  routes.PostRoutes(api)
  routes.AuthRoutes(api)

  api.Get("/", func(c *fiber.Ctx) error {
      return c.SendString("Hello, World!")
  })

  log.Fatal(app.Listen(":3000"))
}
