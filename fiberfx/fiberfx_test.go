package fiberfx_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaskatayama/modfx/fiberfx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type Server struct {
}

var a = []fiberfx.Route{
	{
		Path:   "/hello",
		Method: http.MethodGet,
		Handler: func(ctx *fiber.Ctx) error {
			return ctx.SendString("hello")
		},
	},
}

func TestSimple(t *testing.T) {
	app := fxtest.New(
		t,
		fiberfx.Module(),
		fx.Supply(a),
	)

	app.RequireStart()
	time.Sleep(10 * time.Second)

	app.RequireStop()
}
