package util

func SwapByOrder(i1, i2 int) (min, max int) {
	if i1 < i2 {
		return i1, i2
	}
	return i2, i1
}

func CutString(str string, cutTo int) string {
	if len(str) <= cutTo {
		return str
	}
	if cutTo <= 0 {
		return ""
	}
	return str[:cutTo]
}

func If(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
