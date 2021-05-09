package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
  )
func getInput() ([]string, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	n, _ := strconv.Atoi(text)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	r := regexp.MustCompile("[^\\s]+")
	arrayString :=r.FindAllString(text, -1)
	return arrayString, n
}

func recursiveJump (allPairs []string, size int, position int, numJumps int) (int) {

	if(position+2 >= size){
		if(position+1 <= size){
			return recursiveJump(allPairs, size, position+1, numJumps+1)
		}
		return numJumps-1

	}
	if(allPairs[position+2] == "1"){
		return recursiveJump(allPairs, size, position+1, numJumps+1)
	}
	return recursiveJump(allPairs, size, position+2, numJumps+1)
}

func GetJumps(allPairs []string, n int) (int) {

	return recursiveJump(allPairs, n, 0, 0)
	
}
func main() {

	allPairs, n := getInput();
	numJumps := GetJumps(allPairs, n)
	fmt.Println(numJumps)

}