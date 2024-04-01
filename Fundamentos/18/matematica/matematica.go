package matematica

func Soma[T int | float64](v1, v2 T) T {
	return v1 + v2
}
