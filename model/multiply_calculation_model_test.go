package model

import (
	"testing"
)

func TestNewMultipleCalculation(t *testing.T) {

	sum:=NewMultipleCalculation(1,2)
	gotFirstArg:=sum.FirstArgument
	if(gotFirstArg!=1){
		t.Errorf("want 1 get %d", gotFirstArg)
	}
	gotSecArg:=sum.SecondArgument
	if(gotSecArg!=2){
		t.Errorf("want 2 get %d", gotSecArg)
	}
	gotMultiply:=sum.PrintResult()
	if(gotMultiply!=2){
		t.Errorf("want 2 get %d", gotMultiply)
	}

}
