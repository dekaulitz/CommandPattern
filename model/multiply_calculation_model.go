package model

type MultipleCalculation struct {
	FirstArgument  int
	SecondArgument int
}

func NewMultipleCalculation(firstArgument int, secondArgument int) *MultipleCalculation {
	return &MultipleCalculation{FirstArgument: firstArgument, SecondArgument: secondArgument}
}

func (c MultipleCalculation) PrintResult() int {
	return c.FirstArgument * c.SecondArgument
}
