package binance_futures

import (
	"binanace_coin_trade_system/types"
	"github.com/adshao/go-binance/v2/futures"
)

func CreateShortOrderService(symbol types.BinanceSymbolType, client *futures.Client) *futures.CreateOrderService {

	orderService := client.NewCreateOrderService()
	orderService.Symbol(string(symbol))
	orderService.PositionSide(futures.PositionSideTypeShort)
	orderService.Type(futures.OrderTypeMarket)
	orderService.TimeInForce(futures.TimeInForceTypeGTC)
	orderService.NewClientOrderID("Short RSI")

	return orderService
}
