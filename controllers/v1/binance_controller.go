package v1_contorllers

import (
	"binanace_coin_trade_system/models"
	binance_futures "binanace_coin_trade_system/pkg/lib/binance-futures"
	"binanace_coin_trade_system/types"
	"context"
	"github.com/gofiber/fiber/v2"
)

func GetFuturesChart(c *fiber.Ctx) error {

	getBinanceFuturesChartModel := &models.GetBinanceFuturesChartModel{}

	if err := c.BodyParser(getBinanceFuturesChartModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	futuresCleint := binance_futures.CreateBinanceFuturesClient()

	futuresklines := binance_futures.CreateKlineServiceBySymbol(getBinanceFuturesChartModel.Symbol, getBinanceFuturesChartModel.Interval, futuresCleint)

	response, err := futuresklines.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error()})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success get BTC chart Data",
		"response": fiber.Map{"chart_data": response},
	})
}

func SetFuturesLeverage(c *fiber.Ctx) error {

	setFuturesLeverageModel := &models.SetFuturesLeverageModel{}

	if err := c.BodyParser(setFuturesLeverageModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	futuresCleint := binance_futures.CreateBinanceFuturesClient()
	futuresLeverage := binance_futures.UpdatePositonLeverageBySymbol(setFuturesLeverageModel.Symbol, setFuturesLeverageModel.Leverage, futuresCleint)

	response, err := futuresLeverage.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error()})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success set leverage",
		"response": fiber.Map{"chart_data": response},
	})

}

func SetOrder(c *fiber.Ctx) error {

	futuresCleint := binance_futures.CreateBinanceFuturesClient()
	futuresOrder := binance_futures.CreateShortOrderService(types.BTCUSDT, "0.006", futuresCleint)

	response, err := futuresOrder.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error()})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success set leverage",
		"response": fiber.Map{"chart_data": response},
	})

}
