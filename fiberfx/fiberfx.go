package fiberfx

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"fiberfx",
		fx.Provide(New),
		fx.Invoke(Register),
		fx.Invoke(startApp),
	)
}

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: false,
	})
	return app
}

type Route[T any] struct {
	Path    string
	Method  string
	Handler T
}

type Mounter interface {
	Mount(app *fiber.App)
}

type params struct {
	fx.In
	Mounts []Mounter              `optional:"true"`
	Routes []Route[fiber.Handler] `optional:"true"`
}

func Register(app *fiber.App, p params) {
	if p.Routes != nil {
		for _, route := range p.Routes {
			app.Add(route.Method, route.Path, route.Handler)
		}
	}

	if p.Mounts != nil {
		for _, mounter := range p.Mounts {
			mounter.Mount(app)
		}
	}
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
