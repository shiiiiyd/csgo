package fiber

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"testing"
)

func TestFiber(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world!")
	})
	log.Fatal(app.Listen(":3000"))
}
