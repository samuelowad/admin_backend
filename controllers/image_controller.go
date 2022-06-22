package controllers

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func ImageUpload(c *fiber.Ctx) error {
	path := "./uploads"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			panic(err)
		}
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var filename string

	files := form.File["image"]
	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/uploads/" + filename,
	})

}
