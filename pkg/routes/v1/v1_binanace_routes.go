package v1_routes

import (
	"binanace_coin_trade_system/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func V1BinanceRoute(v1_binanace_route fiber.Router) {

	binanace_route := v1_binanace_route.Group("/binance")

	// Get Coin Future Chart DATA
	binanace_route.Post("future/chart", v1_contorllers.GetFuturesChart)
	binanace_route.Post("future/leverage", v1_contorllers.SetFuturesLeverage)
	binanace_route.Post("future/order/short", v1_contorllers.SetShortOrder)
	binanace_route.Post("future/order/long", v1_contorllers.SetLongOrder)
}
