package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/database"
	"github.com/samuelowad/admin_backend/models"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.Database, &models.User{}, page))

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.Database.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.Database.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.Database.Model(&user).Updates(user)

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.Database.Delete(&user)

	return c.JSON(nil)
}
