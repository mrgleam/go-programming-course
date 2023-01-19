package main

import (
	"final/domain/coupon"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/coupons", coupon.GetCouponHandler(coupon.GetCoupon))
	e.POST("/coupons", coupon.CreateCouponHandler(coupon.CreateCoupon))
	e.GET("/coupons/:id", coupon.GetCouponByIDHandler(coupon.GetCouponByID))
	e.PUT("/coupons/:id", coupon.UpdateCouponByIDHandler(coupon.UpdateCoupon))
	e.DELETE("/coupons/:id", coupon.DeleteCouponHandler(coupon.DeleteCoupon))

	e.Logger.Fatal(e.Start(":1323"))
}
