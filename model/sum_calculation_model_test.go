package model

import (
	"testing"
)

func TestSumCalculation_Input(t *testing.T) {
	sum:=NewSumCalculation(1,2)
	gotFirstArg:=sum.FirstArgument
	if(gotFirstArg!=1){
		t.Errorf("want 1 get %d", gotFirstArg)
	}
	gotSecArg:=sum.SecondArgument
	if(gotSecArg!=2){
		t.Errorf("want 2 get %d", gotSecArg)
	}
	gotSum:=sum.PrintResult()
	if(gotSum!=3){
		t.Errorf("want 3 get %d", gotSum)
	}
}
