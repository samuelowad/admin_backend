package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/database"
	"github.com/samuelowad/admin_backend/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permission []models.Permission

	database.Database.Find(&permission)

	return c.JSON(permission)

}
