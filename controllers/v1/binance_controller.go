package v1_contorllers

import (
	"context"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func FuturesBTCMarkPrice(c *fiber.Ctx) error {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env_data, err := godotenv.Read(".env")
	futures_client := futures.NewClient(env_data["BINANCE_API_KEY"], env_data["BINANCE_SECRET_KEY"])

	futures_client.NewSetServerTimeService().Do(context.Background())

}
