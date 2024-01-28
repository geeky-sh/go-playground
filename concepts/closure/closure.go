package closure

func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func OddInt() func() int {
	first := true
	i := 1
	return func() int {
		if first {
			first = false
			return i
		}
		i += 2
		return i
	}
}

func EvenInt() func() int {
	first := true
	i := 0
	return func() int {
		if first {
			first = false
			return i
		}
		i += 2
		return i
	}
}

func NextNo() func() int {
	first := true
	second := true
	a := 0
	b := 1
	return func() int {
		if first {
			first = false
			return a
		}
		if second {
			second = false
			return b
		}
		b = a + b
		a = b - a
		return b
	}
}
