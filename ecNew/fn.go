package ecNew

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"os/signal"
	"strings"
	"time"
)

func NewEcho() (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	return
}

func Listen(
	e *echo.Echo,
	port string,
	waitShutdown ...time.Duration,
) {
	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	defer stop()
	go func() {
		fmt.Printf("🚀 Server ready at http://localhost%s\n", port)
		if err := e.Start(port); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	var wait = time.Second * 10
	if len(waitShutdown) == 1 {
		wait = waitShutdown[0]
	}

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
