package coupon

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type createCoupon func(Coupon) int

func (fn createCoupon) Create(coupon Coupon) int {
	return fn(coupon)
}

func CreateCouponHandler(gc createCoupon) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon Coupon
		if err := c.Bind(&coupon); err != nil {
			return err
		}
		coupons := gc.Create(coupon)
		return c.JSON(http.StatusOK, coupons)
	}
}
