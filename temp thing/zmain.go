package mane

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var allData = readDb()

func mainee() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/dirList", func(c *fiber.Ctx) error {
		fmt.Print(readDb())
		return c.JSON(allData)
	})

	//TODO make go to DB
	app.Static("/music", "../music", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	fmt.Print("did a thing")
	app.Listen(":9000")
}
