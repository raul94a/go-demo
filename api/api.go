package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{"msg": "I hate Go or maybe a like it 😅"})
	})

	api := app.Group("api")

	v1 := api.Group("v1")

	v1.Get("/users", getUsers)

	err := app.Listen(":8000")

	if err != nil {
		return
	}
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []User{
	{Id: "1", Name: "dorito", Email: "dorito@test", Password: "1234"},
	{Id: "2", Name: "lays", Email: "lays@test", Password: "1234"},
	{Id: "3", Name: "nachos", Email: "nachos@test", Password: "1234"},
	{Id: "4", Name: "ariel", Email: "ariel@test", Password: "1234"},
	{Id: "5", Name: "magi", Email: "magi@test", Password: "1234"},
	{Id: "6", Name: "campurriana", Email: "campurriana@test", Password: "1234"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}
