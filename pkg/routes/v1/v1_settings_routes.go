package v1_routes

import (
	v1_contorllers "binanace_coin_trade_system/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func V1ServerSettingRoute(v1_setting_route fiber.Router) {

	setting_route := v1_setting_route.Group("/setting")

	// Get Coin Future Chart DATA
	setting_route.Get("/bcrypt/password", v1_contorllers.GetServerPassword)
}
