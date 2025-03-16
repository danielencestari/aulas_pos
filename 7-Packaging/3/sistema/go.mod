module github.com/danielencestari/aulas_pos/7-Packagin/3/sistema

go 1.24.0
// replace pra usar localmente, o que é ótimo, mas não é bom pra qdo publicar, pq senão vai da ruim
// mesmo que suba do git vai dar ruim pq não existe url relativa no git
// não é uma boa usar essa gambiarra, por isso deixei ela comentada, para ver a proxima solução
// replace github.com/danielencestari/aulas_pos/7-Packaging/3/math => ../math

// require github.com/danielencestari/aulas_pos/7-Packaging/3/math v0.0.0-00010101000000-000000000000
