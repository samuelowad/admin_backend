package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/database"
	"github.com/samuelowad/admin_backend/models"
	"strconv"
)

type RoleCreateDTO struct {
	Name        string
	permissions []string
}

func AllRole(c *fiber.Ctx) error {
	var roles []models.Role

	database.Database.Find(&roles)

	return c.JSON(roles)

}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	database.Database.Create(&role)

	return c.JSON(role)

}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.Database.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}
	database.Database.Table("role_permissions").Where("role_id = ?", id).Delete(&result)

	role := models.Role{
		Id:         uint(id),
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}
	database.Database.Model(&role).Updates(role)

	return c.JSON(role)

}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.Database.Delete(&role)

	return c.JSON(nil)
}
