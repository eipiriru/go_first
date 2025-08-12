package helpers

func Ternary(cond bool, a, b any) any {
	if cond {
		return a
	}
	return b
}
