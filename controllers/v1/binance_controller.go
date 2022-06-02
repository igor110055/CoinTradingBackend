package v1_contorllers

import (
	"binanace_coin_trade_system/models"
	"binanace_coin_trade_system/pkg/lib/binance-futures"
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
		"response": fiber.Map{"leverage": response},
	})

}

func SetShortOrder(c *fiber.Ctx) error {

	setShortOrderModel := &models.SetShortOrderModel{}

	if err := c.BodyParser(setShortOrderModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	stopPriceTP := setShortOrderModel.StopPriceTP
	stopPriceSL := setShortOrderModel.StopPriceSL

	futuresClient := binance_futures.CreateBinanceFuturesClient()
	futuresOrder := binance_futures.CreateShortOrderService(setShortOrderModel.Symbol, "0.005", futuresClient)
	futuresOrderTP := binance_futures.CreateShortTP(setShortOrderModel.Symbol, &stopPriceTP, futuresClient)
	futuresOrderSL := binance_futures.CreateShortSL(setShortOrderModel.Symbol, &stopPriceSL, futuresClient)

	binance_futures.CancelAllOpenOrder(setShortOrderModel.Symbol, futuresClient)

	resShort, err := futuresOrder.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short Error"})
	}

	resShortTP, err := futuresOrderTP.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short TP Error"})
	}

	resShortSL, err := futuresOrderSL.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short SL Error",
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success set leverage",
		"response": fiber.Map{"resShort": resShort, "resShortTP": resShortTP, "resShortSL": resShortSL},
	})

}

func SetLongOrder(c *fiber.Ctx) error {

	setLongOrderModel := &models.SetLongOrderModel{}

	if err := c.BodyParser(setLongOrderModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	stopPriceTP := setLongOrderModel.StopPriceTP
	stopPriceSL := setLongOrderModel.StopPriceSL

	futuresClient := binance_futures.CreateBinanceFuturesClient()
	futuresOrder := binance_futures.CreateLongOrderService(setLongOrderModel.Symbol, "0.005", futuresClient)
	futuresOrderTP := binance_futures.CreateLongTP(setLongOrderModel.Symbol, &stopPriceTP, futuresClient)
	futuresOrderSL := binance_futures.CreateLongSL(setLongOrderModel.Symbol, &stopPriceSL, futuresClient)

	binance_futures.CancelAllOpenOrder(setLongOrderModel.Symbol, futuresClient)

	resShort, err := futuresOrder.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short Error"})
	}

	resShortTP, err := futuresOrderTP.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short TP Error"})
	}

	resShortSL, err := futuresOrderSL.Do(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":       true,
			"msg":         err.Error(),
			"description": "Short SL Error",
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "success set leverage",
		"response": fiber.Map{"resShort": resShort, "resShortTP": resShortTP, "resShortSL": resShortSL},
	})

}
