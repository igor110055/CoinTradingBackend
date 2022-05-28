package binance_futures

import (
	"context"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/joho/godotenv"
	"log"
)

func CreateBinanceFuturesClient() *futures.Client {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envData, err := godotenv.Read(".env")
	futuresClient := futures.NewClient(envData["BINANCE_API_KEY"], envData["BINANCE_SECRET_KEY"])

	_, err = futuresClient.NewSetServerTimeService().Do(context.Background())

	return futuresClient
}
