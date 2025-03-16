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

// como testar varias condições ao mesmo tempo
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500, 5.0},
		{1000, 10.0},
		{1500, 10.0},
		// {1500, 15.0}, // erro, que só aparece se rodar com um go test -v
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expected {
			t.Errorf("Expected %f, got %f", item.expected, result)
		}
	}

}
