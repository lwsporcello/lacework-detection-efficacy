package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		log.Infof("Hit Listener Beacon URL with body %s", c.Body())
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Get("/:filename", func(c *fiber.Ctx) error {
		log.Infof("File download URL with filename %s", c.Params("filename"))

		filename := c.Params("filename")
		if filename == "install-demo-1.sh" {
			dir, _ := os.Getwd()
			return c.SendFile(dir + "/bin/" + filename)
		}
		return c.SendStatus(fiber.StatusNotFound)
	})
	_ = app.Listen(":8081")
}
