package main

import (
	"github.com/dileep9490/todoapp/Backend/database"
	"github.com/dileep9490/todoapp/Backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Post("/auth/signup", handlers.SignUP)
	app.Post("/auth/login", handlers.Login)
	app.Post("/todo/create", handlers.CreateTodo)
	app.Get("/todo/all/:user_uid", handlers.GetAllTodos)
	app.Get("/todo/:id", handlers.GetTodoById)
	app.Put("/todo/update", handlers.UpdataTodo)
	app.Delete("/todo/delete/:id", handlers.DeleteTodo)
	app.Listen(":8080")
}
