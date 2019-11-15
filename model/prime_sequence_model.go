package model

import (
	"strconv"
	"strings"
)

type PrimeSequnce struct {
	Input int
}

func NewPrimeChar(input int) *PrimeSequnce {
	return &PrimeSequnce{Input: input}
}

func (p *PrimeSequnce) GenerateSquence(x int) int {
	if x <= 1 {
		return 0
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return 0
		}
	}
	return 1
}

func (p *PrimeSequnce) PrintResult() string {
	result := []string{}
	sum, max := 0, 0
	for max < p.Input {
		if p.GenerateSquence(sum) == 1 {
			result = append(result, strconv.Itoa(sum))
			max++
		}
		sum++
	}
	return strings.Join(result, ",")
}
