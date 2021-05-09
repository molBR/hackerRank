


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  )
func getInput() (string, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	str := strings.Replace(text, "\n", "", -1)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	n, _ := strconv.Atoi(text)
	
	return str, n
}


func ReturnFreqA(str string, n int) (int) {
	numberOfOcurrences := 0
	for i:=0; i < len(str); i++{
		if(string(str[i]) == "a") {
			numberOfOcurrences++
		}
	}
	numberOfOcurrences = numberOfOcurrences*(n/len(str))
	restOfSweep := n%len(str)
	for i:=0; i<restOfSweep; i++ {
		if(string(str[i]) == "a"){
			numberOfOcurrences++
		}
	}
	return numberOfOcurrences


}


func main() {

	str, n := getInput();
	fmt.Println(ReturnFreqA(str, n))
}