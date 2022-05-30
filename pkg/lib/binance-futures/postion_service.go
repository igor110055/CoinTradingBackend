package binance_futures

import (
	"binanace_coin_trade_system/types"
	"github.com/adshao/go-binance/v2/futures"
)

func UpdatePositonLeverageBySymbol(symbol types.BinanceSymbolType, leverage int, client *futures.Client) *futures.ChangeLeverageService {

	leverageService := client.NewChangeLeverageService()
	leverageService.Symbol(string(symbol))
	leverageService.Leverage(leverage)

	return leverageService
}
