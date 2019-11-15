package model

import (
	"strconv"
	"strings"
)

type FibonaciSquence struct {
	Input int
}

func NewFibonaciSquence(input int) *FibonaciSquence {
	return &FibonaciSquence{Input: input}
}

func (f *FibonaciSquence) GenerateSquence(x int) int {
	if x <= 1 {
		return x
	}
	return f.GenerateSquence(x-1) + f.GenerateSquence(x-2)
}
func (f *FibonaciSquence) PrintResult() string {
	result := []string{}
	for i := 0; i < f.Input; i++ {
		result = append(result, strconv.Itoa(f.GenerateSquence(i)))
	}
	return strings.Join(result, ",")
}
