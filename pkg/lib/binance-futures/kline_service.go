package binance_futures

import (
	"binanace_coin_trade_system/types"
	"github.com/adshao/go-binance/v2/futures"
)

func CreateKlineServiceBySymbol(symbol types.BinanceSymbolType, interval types.BinanceIntervalType, client *futures.Client) *futures.KlinesService {

	klineService := client.NewKlinesService()
	klineService.Interval(string(interval))
	klineService.Symbol(string(symbol))

	return klineService
}
