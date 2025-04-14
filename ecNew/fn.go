package ecNew

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	return
}
