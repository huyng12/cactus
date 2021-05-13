package handler

import "github.com/gofiber/fiber/v2"

type HealthCheckResponse struct {
	Alive bool `json:"alive"`
}

func HealthCheck(ctx *fiber.Ctx) error {
	status := HealthCheckResponse{Alive: true}
	return ctx.JSON(status)
}
