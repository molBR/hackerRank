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

func GetPairs(allPairs []string, n int) (int) {
	mapPairs:=make(map[int]int)
	for i:= 0; i<n;i++{
		intPair,_:=strconv.Atoi(allPairs[i])
		mapPairs[intPair] = mapPairs[intPair] + 1
	}
	var pairs int = 0
	for _, element := range mapPairs {
		pairs = pairs + element / 2
	}
	return pairs
	
}

func main() {

	allPairs, n := getInput();
	pairs:=GetPairs(allPairs, n)
	fmt.Println(pairs)
}