package main

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, val := range m {
		soma += val
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mp := map[string]int{"A": 1, "b": 2, "C": 3, "e": 4, "F": 5}
	mp2 := map[string]float64{"A": 1.5, "b": 2.6, "C": 3.0, "e": 4.0, "F": 5}
	println(Soma(mp))
	println(Soma(mp2))
	print(Compara(10, 10))
}
