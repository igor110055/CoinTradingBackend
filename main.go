package main

import (
	"binanace_coin_trade_system/pkg/config"
	"binanace_coin_trade_system/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.FiberConfig()

	app := fiber.New(config)

	//routes
	routes.RoutesIndex(app)

	app.Listen(":3000")
}
