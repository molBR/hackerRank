package main
import "testing"


func TestGetJumps(t *testing.T){

	pairStrings := []string{"0", "0", "1", "0", "0", "1", "0"}
	returnValue := GetJumps(pairStrings, 7)
	if (!(returnValue == 4)){
		t.Error()
	}

	pairStrings = []string{"0", "0", "0", "0", "1", "0"}
	returnValue = GetJumps(pairStrings, 6)
	if (!(returnValue == 3)){
		t.Error()
	}
}