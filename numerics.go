package gopromax

type NumericType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func MaxNumber[T NumericType](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MinNumber[T NumericType](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func ClapNumber[T NumericType](n, left, right T) T {
	if n < left {
		return left
	} else if n > right {
		return right
	} else {
		return n
	}
}
