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

// para ver a cobertura dos testes, executar:
// go test -coverprofile=coverage.out
// para ver onde não foi coberto, tem que executar:
// go tool cover -html=coverage.out
// ai ele abre o browser com o relatório de cobertura, e o que está em vermelho não foi coberto

// teste de benchmark usa o b do tipo testing.B
// para rodar o benchmark, executar:
// go test -bench=.
// na aula o resultado do prof foi diferente, pois mostrou mais detalhes da operacao
// para ver mais detalhes, executar
// go test -bench=. -benchmem
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

// rodando: go test -bench=. -run=ˆ#
// da pra ver a diferença do tempo de execução entre os dois testes
// o temporizador incluido na func CalculateTax2 faz com que ela seja mais lenta
// se rodar go test -bench=. -run=ˆ# -count=10, ele roda 10 vezes
// se rodar go test -bench=. -run=ˆ# -count=10 -benchtime=10s, ele roda 10 vezes e cada vez dura 10s
func BenchmarkCalculate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

// Fuzzing - testes de mutação
// o fuzz é tipo uma rede neural que tenta encontrar falhas no código
// vc tem que "treinar" um pouco o fuzz, para ele entender o que é um erro
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}
