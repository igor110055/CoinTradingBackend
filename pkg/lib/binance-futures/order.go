package binance_futures

import (
	"binanace_coin_trade_system/types"
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
	orderService.NewClientOrderID("ShortRSI")

	return orderService
}
