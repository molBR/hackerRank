package main
import "testing"


func TestGetJumps(t *testing.T){

	str := "abcac"
	returnValue :=  ReturnFreqA(str, 10)
	if (!(returnValue == 4)){
		t.Error()
	}

}