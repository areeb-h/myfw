package framework

import "github.com/gofiber/fiber/v2"

// RequestBody automatically parses JSON request body
func RequestBody[T any](c *fiber.Ctx) (T, error) {
	var body T
	err := c.BodyParser(&body)
	return body, err
}
