package v1_routes

import (
	"binanace_coin_trade_system/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func V1BinanceRoute(v1_binanace_route fiber.Router) {

	binanace_route := v1_binanace_route.Group("/binance")

	// Get BTC Coin mark_price
	binanace_route.Get("future/btc/chart", v1_contorllers.FuturesBTCChart)
}
