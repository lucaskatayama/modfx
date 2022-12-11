package fiberfx_test

import (
	"testing"
	"time"

	"github.com/lucaskatayama/modfx/fiberfx"
	"go.uber.org/fx/fxtest"
)

func TestSimple(t *testing.T) {
	app := fxtest.New(
		t,
		fiberfx.Module(),
	)

	app.RequireStart()
	time.Sleep(10 * time.Second)

	app.RequireStop()
}
