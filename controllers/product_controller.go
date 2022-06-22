package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/database"
	"github.com/samuelowad/admin_backend/models"
	"strconv"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.Database, &models.Product{}, page))

}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.Database.Create(&product)

	return c.JSON(product)

}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.Database.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.Database.Model(&product).Updates(product)

	return c.JSON(product)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.Database.Delete(&product)

	return c.JSON(nil)
}
