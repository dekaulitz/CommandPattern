package main

import (
	_interface "awesomeProject/interface"
	"awesomeProject/model"
	"fmt"
)

func main() {
	multiple := model.NewMultipleCalculation(1, 2)
	sum := model.NewSumCalculation(1, 2)
	command := model.NewCommandCalculation()
	err := command.SetCommand(multiple, _interface.MultiplySeq)
	if (err != nil) {
		fmt.Println(err)
	}
	command.Execute()
	err = command.SetCommand(sum, _interface.SumSeq)
	if (err != nil) {
		fmt.Println(err)
	}
	command.Execute()
	commandSequence := model.NewCommandSequence()
	fibonacu := model.NewFibonaciSquence(4)
	err = commandSequence.SetCommand(fibonacu, _interface.FibonaciSeq)
	if (err != nil) {
		fmt.Println(err)
	}
	commandSequence.Execute()
	prime := model.NewPrimeChar(4)
	err = commandSequence.SetCommand(prime, _interface.PrimeSeq)
	if (err != nil) {
		fmt.Println(err)
	}
	commandSequence.Execute()

}
