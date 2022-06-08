package functions

func CalculatePrice(price float32, quantity int, discount float32) float32 {
	if discount > 0 {
		return price * (1 - (discount / 100)) * float32(quantity)
	}
	return price * float32(quantity)
}
