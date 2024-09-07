package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) int64 {
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		log.Fatal("userID not found somehow")
	}
	return userID
}
