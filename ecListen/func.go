package ecListen

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"strings"
	"time"
)

func Listen(
	e *echo.Echo,
	port string,
	waitShutdown ...time.Duration,
) {
	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	var ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt)

	defer stop()
	go func() {
		fmt.Printf("🚀 HTTP server ready at http://localhost%s\n", port)

		var err error
		if err = e.Start(port); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	var wait = time.Second * 10
	if len(waitShutdown) == 1 {
		wait = waitShutdown[0]
	}

	var cancel func()
	ctx, cancel = context.WithTimeout(context.Background(), wait)
	defer cancel()

	var err error
	if err = e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
