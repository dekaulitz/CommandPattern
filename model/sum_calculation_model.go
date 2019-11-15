package model

type SumCalculation struct {
	FirstArgument  int
	SecondArgument int
}

func NewSumCalculation(firstArgument int, secondArgument int) *SumCalculation {
	return &SumCalculation{FirstArgument: firstArgument, SecondArgument: secondArgument}
}

func (c SumCalculation) PrintResult() int {
	return c.FirstArgument + c.SecondArgument
}
