package math

type Number interface {
	int64 | int32 | float64 | float32
}

func Max[T Number](num1, num2 T) T {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

func Min[T Number](num1, num2 T) T {
	if num1 <= num2 {
		return num1
	} else {
		return num2
	}
}
