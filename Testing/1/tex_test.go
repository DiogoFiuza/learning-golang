package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type CalcTax struct {
		amount, expect float64
	}

	table := []CalcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f, got %f", item.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(1000.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Recived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Recived %f but expected 20", result)
		}
	})
}

// Commands to run the tests:
// go test: run all tests in the current directory
// go test -v: run all tests in the current directory with verbose output
// go test -coverprofile=coverage.out: generate a coverage profile
// go tool cover -html=coverage.out: view the coverage profile in a web browser
// go test -bench=.: run all benchmarks in the current directory
// go test --fuzz=. : run all fuzz tests in the current directory
