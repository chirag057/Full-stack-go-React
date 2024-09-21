package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	fmt.Println("Hello Worldss")
	app:=fiber.New()

	app.Get("/",func(c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg" :"Hello World"})
	})
	
	log.Fatal(app.Listen(":4000"))   // To check the errors this is used

	
}
