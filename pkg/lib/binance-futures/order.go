package binance_futures

import (
	"binanace_coin_trade_system/types"
	"context"
	"github.com/adshao/go-binance/v2/futures"
)

func CreateShortOrderService(symbol types.BinanceSymbolType, quantity string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeSell)
	orderService.PositionSide(futures.PositionSideTypeShort)
	// Quantatoy = markprice * quantity / leverage = margin
	orderService.Quantity(quantity)
	orderService.Type(futures.OrderTypeMarket)
	orderService.NewClientOrderID("Short")

	return orderService
}

func CreateShortTP(symbol types.BinanceSymbolType, stopPriceTP *string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeBuy)
	orderService.PositionSide(futures.PositionSideTypeShort)
	orderService.Type(futures.OrderTypeTakeProfitMarket)
	orderService.ClosePosition(true)
	orderService.StopPrice(*stopPriceTP)
	orderService.NewClientOrderID("TPShort")

	return orderService
}

func CreateShortSL(symbol types.BinanceSymbolType, stopPriceSL *string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeBuy)
	orderService.PositionSide(futures.PositionSideTypeShort)
	orderService.Type(futures.OrderTypeStopMarket)
	orderService.ClosePosition(true)
	orderService.StopPrice(*stopPriceSL)
	orderService.NewClientOrderID("SLShort")

	return orderService
}

func CancelAllOpenOrder(symbol types.BinanceSymbolType, client *futures.Client) {

	cancleAllOpenOrderService := client.NewCancelAllOpenOrdersService()
	cancleAllOpenOrderService.Symbol(string(symbol))

	cancleAllOpenOrderService.Do(context.Background())

}

func CreateLongOrderService(symbol types.BinanceSymbolType, quantity string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeBuy)
	orderService.PositionSide(futures.PositionSideTypeLong)
	// Quantatoy = markprice * quantity / leverage = margin
	orderService.Quantity(quantity)
	orderService.Type(futures.OrderTypeMarket)
	orderService.NewClientOrderID("Long")

	return orderService
}

func CreateLongTP(symbol types.BinanceSymbolType, stopPriceTP *string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeSell)
	orderService.PositionSide(futures.PositionSideTypeLong)
	orderService.Type(futures.OrderTypeTakeProfitMarket)
	orderService.ClosePosition(true)
	orderService.StopPrice(*stopPriceTP)
	orderService.NewClientOrderID("TPLong")

	return orderService
}

func CreateLongSL(symbol types.BinanceSymbolType, stopPriceSL *string, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.Side(futures.SideTypeSell)
	orderService.PositionSide(futures.PositionSideTypeLong)
	orderService.Type(futures.OrderTypeStopMarket)
	orderService.ClosePosition(true)
	orderService.StopPrice(*stopPriceSL)
	orderService.NewClientOrderID("SLLong")

	return orderService
}
