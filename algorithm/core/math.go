package core

type MyInt interface {
	int | int8 | int16 | int32 | int64
}

func Max[T MyInt](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T MyInt](a, b T) T {
	if a > b {
		return b
	}

	return a
}

func Abs[T MyInt](a T) T {
	if a > 0 {
		return a
	}
	return -a
}
