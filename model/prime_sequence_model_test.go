package model

import (
	"testing"
)

func TestPrimeChar_GenerateSquence(t *testing.T) {
	f := NewPrimeChar(4)
	gotInput:=f.Input
	if(gotInput!=4){
		t.Errorf("want 4 get %d", gotInput)
	}
	gotInput=f.GenerateSquence(5)
	if(gotInput!=1){
		t.Errorf("want 1 get %d", gotInput)
	}
}
func TestPrimeSequnce_PrintResult(t *testing.T) {
	f := NewPrimeChar(4)
	gotInput:=f.Input
	if(gotInput!=4){
		t.Errorf("want 4 get %d", gotInput)
	}
	gotResult:=f.PrintResult()
	if(gotResult!="2,3,5,7"){
		t.Errorf("want 2,3,5,7 get %s", gotResult)
	}
}
