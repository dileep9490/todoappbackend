package handlers

import (
	"net/http"

	"github.com/dileep9490/todoapp/Backend/database"
	"github.com/dileep9490/todoapp/Backend/models"
	"github.com/dileep9490/todoapp/Backend/utils/types"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	data := new(types.TodoType)
	db := database.DB
	todo := new(models.Todo)
	if err := c.BodyParser(data); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "error in creating todo",
		})
	}

	todo.User_uid = data.User_uid
	todo.Title = data.Title

	if err := db.Create(&todo).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": todo,
	})
}

func GetAllTodos(c *fiber.Ctx) error {
	var alltodos []models.Todo
	user_uid := c.Params("user_uid")
	db := database.DB

	db.Model(&models.Todo{}).Find(&alltodos, "user_uid=?", user_uid)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": alltodos,
	})
}

func GetTodoById(c *fiber.Ctx) error {

	id := c.Params("id")

	todo := new(models.Todo)

	db := database.DB

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "todo not found",
		})
	}

	return c.Status(http.StatusOK).JSON(todo)

}

func UpdataTodo(c *fiber.Ctx) error {
	data := new(types.TodoUpdate)
	todo := new(models.Todo)
	db := database.DB

	if err := c.BodyParser(data); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	if err := db.Where("id = ?", data.ID).First(&todo).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err,
		})
	}

	todo.Completed = data.Completed
	todo.Title = data.Title
	println(todo)
	if err := db.Save(&todo).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(todo)

}


func DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB
	todo:= new(models.Todo)

	if err:=db.Where("id = ?",id).Unscoped().Delete(&todo).Error;err!=nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":"couldn't delete Todo",
		})
	}

	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status" : "deleted succesfully",
		"id" : id,
	})
	
}