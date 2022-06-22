package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/samuelowad/admin_backend/database"
	"github.com/samuelowad/admin_backend/models"
	"github.com/samuelowad/admin_backend/util"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func Register(c *fiber.Ctx) error {
	data, err := getRequestBody(c)

	if err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid password and confirmation password",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}
	user.SetPassword(data["password"])

	err = database.Database.Create(&user).Error
	if err != nil {
		fmt.Println(err)

	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data, err := getRequestBody(c)

	if err != nil {
		return err
	}

	var user models.User

	database.Database.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return badRequest(c, "incorrect login ")
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		return badRequest(c, "error logining in ")
	}
	cookie := fiber.Cookie{
		Name:     "go-back",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "login successful"})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "go-back",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON("logout successful")
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("go-back")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.Database.Where("id=?", id).First(&user)

	return c.JSON(user)

}

func UpdateInfo(c *fiber.Ctx) error {
	data, _ := getRequestBody(c)

	cookie := c.Cookies("go-back")

	id, _ := (util.ParseJwt(cookie))

	userId, _ := strconv.Atoi(id)

	var user = models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.Database.Model(&user).Updates(data)

	return c.JSON(user)

}

func getRequestBody(c *fiber.Ctx) (map[string]string, error) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return data, err
	}

	return data, nil

}

//badRequest returns 400 error
func badRequest(c *fiber.Ctx, message string) error {
	c.Status(400)
	return c.JSON(fiber.Map{
		"message": message,
	})
}
