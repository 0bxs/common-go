package wallet

func Integral(price uint32, ratio uint32) uint32 {
	return (price / 1000 * ratio) / 1000 * 1000
}
