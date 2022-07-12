package controllers

import (
  "github.com/gofiber/fiber/v2"
  "github.com/smc181002/jwt-auth/middleware"
)

type Payload struct {
  Name string `json:"name"` 
  Password string `json:"password"` 
}

func AuthenticateUser(c *fiber.Ctx) error {
  var payload Payload
  if err := c.BodyParser(&payload); err != nil {
    return fiber.ErrBadRequest
  }

  if !middleware.CheckUser(payload.Name, payload.Password) {
    return fiber.ErrBadRequest
  }

  tokenVal, err := middleware.GenJWT(payload.Name) 

  if (err != nil) {
    return fiber.ErrBadRequest
  }

  return c.JSON(fiber.Map{
    "token": tokenVal,
  })
}
