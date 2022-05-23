package v1_routes

import "github.com/gofiber/fiber/v2"

func V1BinanceRoute(v1_binance_route fiber.Router) {

	binance_route := v1_binance_route.Group("/binance")

	binance_route.Get("check")
}
