package _interface


type BaseCalculation interface {
	PrintResult() int
}

type BaseSequenceChar interface {
	GenerateSquence(x int) int
	PrintResult()string
}
type TypeCalculation string

const (
	FibonaciSeq TypeCalculation = "FIBONACI"
	PrimeSeq    TypeCalculation = "PRIME"
	MultiplySeq TypeCalculation = "MULTIPLY"
	SumSeq      TypeCalculation = "SUM"
)

type BaseCommand interface {
	SetCommand(any interface{}, command TypeCalculation) error
	Execute()
}
