package v1_contorllers

import (
	custom_bcrypt "binanace_coin_trade_system/pkg/lib/bcrypt"
	"github.com/gofiber/fiber/v2"
)

func GetServerPassword(c *fiber.Ctx) error {
	custom_bcrypt.CreateBcryptPassword()

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Set Complete",
	})
}
