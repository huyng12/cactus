package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyng12/cactus/internal/wallet"
)

type CheckBalanceResponse struct {
	Balance int `json:"balance"`
}

func CheckBalance(w wallet.Walleter) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		currentBalance := w.GetBalance()
		return ctx.JSON(CheckBalanceResponse{Balance: currentBalance})
	}
}
