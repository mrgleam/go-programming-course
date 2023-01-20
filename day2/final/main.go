package main

import (
	"final/domain/coupon"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/coupons", coupon.GetCouponHandler(coupon.GetCoupon))
	e.POST("/coupons", coupon.CreateCouponHandler(coupon.CreateCoupon))
	e.GET("/coupons/:id", coupon.GetCouponByIDHandler(coupon.GetCouponByID))
	e.PUT("/coupons/:id", coupon.UpdateCouponByIDHandler(coupon.UpdateCoupon))
	e.DELETE("/coupons/:id", coupon.DeleteCouponHandler(coupon.DeleteCoupon))

	e.Logger.Fatal(e.Start(":1323"))
}
