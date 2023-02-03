package app

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ResponsePayload struct {
	IP        string `json:"ip"`
	Uptime    string `json:"uptime"`
	StartedAt string `json:"started_at"`
}

func uptime(initial time.Time) time.Duration {
	return time.Since(initial)
}

func Start() {
	startedAt := time.Now()

	logger, _ := zap.NewProduction()

	defer logger.Sync()

	logger.Info("Application started...",
		// Structured context as strongly typed Field values.
		zap.Time("startedAt", startedAt),
	)

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		logger.Info(fmt.Sprintf("Request handled %v", ctx.IP()),
			zap.String("url", ctx.BaseURL()),
			zap.String("ip", ctx.IP()),
		)

		return ctx.JSON(ResponsePayload{
			IP:        ctx.IP(),
			Uptime:    uptime(startedAt).String(),
			StartedAt: startedAt.Format(time.RFC3339),
		})
	})
	app.Listen(":3000")
}
