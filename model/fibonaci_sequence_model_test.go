package model

import (
	"testing"
)

func TestNewFibonaciSquence(t *testing.T) {
	f:=NewFibonaciSquence(4)
	gotInput:=f.Input
	if(gotInput!=4){
		t.Errorf("want 4 get %d", gotInput)
	}
	gotInput=f.GenerateSquence(1)
	if(gotInput!=1){
		t.Errorf("want 1 get %d", gotInput)
	}
}

func TestFibonaciSquence_GenerateSquence(t *testing.T) {
	f:=NewFibonaciSquence(4)
	gotInput:=f.Input
	if(gotInput!=4){
		t.Errorf("want 4 get %d", gotInput)
	}
	gotStringInput:=f.PrintResult()
	if(gotStringInput!= "0,1,1,2"){
		t.Errorf("want 0,1,1,2 get %s", gotStringInput)
	}
}