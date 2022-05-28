package v1_contorllers

import (
	binance_futures "binanace_coin_trade_system/pkg/lib/binance-futures"
	"binanace_coin_trade_system/types"
	"context"
	"github.com/gofiber/fiber/v2"
)

func FuturesBTCChart(c *fiber.Ctx) error {

	futuresCleint := binance_futures.CreateBinanceFuturesClient()

	futuresklines := binance_futures.CreateKlineServiceBySymbol(types.BTCUSDT, types.M15, futuresCleint)

	response, err := futuresklines.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error()})
	}

	return c.JSON(fiber.Map{
		"error":      false,
		"msg":        "success get BTC chart Data",
		"chart_data": fiber.Map{"response": response},
	})
}
