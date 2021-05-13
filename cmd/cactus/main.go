package main

import (
	"github.com/huyng12/cactus/internal/handler"
	"github.com/huyng12/cactus/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Intitate a wallet for doing transactions
	myWallet := wallet.New()

	// Routes
	app.Get("/health_check", handler.HealthCheck)
	app.Get("/balance", handler.CheckBalance(myWallet))

	_ = app.Listen(":3000")
}
