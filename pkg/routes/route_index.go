package routes

import (
	"binanace_coin_trade_system/pkg/routes/v1"
	"github.com/gofiber/fiber/v2"
)

func RoutesIndex(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1_routes.V1BinanceRoute(v1)
}
