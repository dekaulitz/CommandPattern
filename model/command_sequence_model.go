package model

import (
	"awesomeProject/error_handler"
	_interface "awesomeProject/interface"
)

type CommandSequence struct {
	Sequence _interface.BaseSequenceChar
}

func NewCommandSequence() *CommandSequence {
	return &CommandSequence{}
}

func (c *CommandSequence) SetCommand(any _interface.BaseSequenceChar, command _interface.TypeCalculation) error {
	switch command {
	case _interface.PrimeSeq:
		c.Sequence = any
		break
	case _interface.FibonaciSeq:
		c.Sequence = any
		break
	default:
		return error_handler.NewAppException("invalid command")
	}
	return nil
}

func (c *CommandSequence) Execute() () {
	println(c.Sequence.PrintResult())
}
