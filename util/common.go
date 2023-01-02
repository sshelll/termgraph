package util

func SwapByOrder(i1, i2 int) (min, max int) {
	if i1 < i2 {
		return i1, i2
	}
	return i2, i1
}
