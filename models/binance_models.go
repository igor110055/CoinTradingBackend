package models

import "binanace_coin_trade_system/types"

type GetBinanceFuturesChartModel struct {
	Symbol   types.BinanceSymbolType   `json:"symbol" validate:"required,lte=255"`
	Interval types.BinanceIntervalType `json:"interval" validate:"required,lte=255"`
}

type SetFuturesLeverageModel struct {
	Symbol   types.BinanceSymbolType `json:"symbol" validate:"required,lte=255"`
	Leverage int                     `json:"leverage" validate:"required,lte=255"`
}
