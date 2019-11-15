package model

import (
	"awesomeProject/error_handler"
	_interface "awesomeProject/interface"
)

type CommandCalculation struct {
	Calc _interface.BaseCalculation
}

func NewCommandCalculation() *CommandCalculation {
	return &CommandCalculation{}
}

func (c *CommandCalculation) SetCommand(any _interface.BaseCalculation, command _interface.TypeCalculation) error {
	switch command {
	case _interface.MultiplySeq:
		c.Calc = any
		break
	case _interface.SumSeq:
		c.Calc = any
		break
	default:
		return error_handler.NewAppException("invalid command")
	}
	return nil
}

func (c *CommandCalculation) Execute() () {
	println(c.Calc.PrintResult())
}
