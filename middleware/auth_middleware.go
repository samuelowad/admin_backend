package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/util"
)

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("go-back")
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
