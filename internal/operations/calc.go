package operations

// Operação matemática de exemplo de exemplo

type Calc struct {
	A float64
	B float64
}

// Soma
func (c Calc) Sum() float64 {
	return c.A + c.B
}

// Construtor
func NewCalc() Calc {
	return Calc{}
}
