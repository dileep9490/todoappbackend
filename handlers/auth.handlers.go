package handlers

import (
	"net/http"

	"github.com/dileep9490/todoapp/Backend/database"
	"github.com/dileep9490/todoapp/Backend/models"
	"github.com/dileep9490/todoapp/Backend/utils"
	"github.com/dileep9490/todoapp/Backend/utils/types"
	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

func SignUP(c *fiber.Ctx) error {
	db := database.DB
	data := new(types.SignUpType)
	user := new(models.User)

	if err := c.BodyParser(data); err != nil {
		c.JSON(fiber.Map{
			"error": err,
		})
	}

	hashpassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot Hash password",
		})
	} else {
		user.Password = hashpassword
	}

	user.Email = data.Email
	user.Name = data.Name
	user.ID = uuid.New()
	if err := db.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": user})
}

func Login(c *fiber.Ctx) error {
	data := new(types.LoginType)
	db := database.DB
	user := new(models.User)

	if err := c.BodyParser(data); err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	db.First(&user, "email=?", data.Email)

	if !(utils.ComparePassword(user.Password, data.Password)) {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "credentials don't match",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"apiKey": user.ID,
	})
}
