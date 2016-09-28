package pkg

type CouponAgent service
type Coupon struct{}

func (*CouponAgent) Get() *Coupon {

	coupon := &Coupon{}
	return coupon
}
