package main
import "testing"


func TestGetPairs(t *testing.T){

	pairStrings := []string{"10", "20", "20", "10", "10", "30", "50", "10", "20"}
	returnValue := GetPairs(pairStrings, 9)
	if (!(returnValue == 3)){
		t.Error()
	}
}