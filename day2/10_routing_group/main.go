package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/coupons/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/coupons/:id")
	})

	g := e.Group("/admin")

	g.GET("/coupons/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/admin/coupons/:id")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
