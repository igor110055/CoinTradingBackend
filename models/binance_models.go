package models

import (
	"binanace_coin_trade_system/types"
)

type GetBinanceFuturesChartModel struct {
	Symbol   types.BinanceSymbolType   `json:"symbol" validate:"required,lte=255"`
	Interval types.BinanceIntervalType `json:"interval" validate:"required,lte=255"`
	Password string                    `json:"password" validate:"required,lte=255"`
}

type SetFuturesLeverageModel struct {
	Symbol   types.BinanceSymbolType `json:"symbol" validate:"required,lte=255"`
	Leverage int                     `json:"leverage" validate:"required,lte=255"`
	Password string                  `json:"password" validate:"required,lte=255"`
}

type SetShortOrderModel struct {
	Symbol      types.BinanceSymbolType `json:"symbol" validate:"required,lte=255"`
	StopPriceTP string                  `json:"stop_price_tp" validate:"required,lte=255"`
	StopPriceSL string                  `json:"stop_price_sl" validate:"required,lte=255"`
	Password    string                  `json:"password" validate:"required,lte=255"`
}

type SetLongOrderModel struct {
	Symbol      types.BinanceSymbolType `json:"symbol" validate:"required,lte=255"`
	StopPriceTP string                  `json:"stop_price_tp" validate:"required,lte=255"`
	StopPriceSL string                  `json:"stop_price_sl" validate:"required,lte=255"`
	Password    string                  `json:"password" validate:"required,lte=255"`
}
