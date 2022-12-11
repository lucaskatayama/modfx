package fiberfx

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"fiberfx",
		fx.Provide(createApp),
		fx.Invoke(startApp),
	)
}

func createApp() *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: false,
	})
	return app
}

func startApp(lf fx.Lifecycle, app *fiber.App) {
	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(":8080"); err != nil {
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
