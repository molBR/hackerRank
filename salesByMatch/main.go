package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
  )
  

func main() {


	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	n, _ := strconv.Atoi(text)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	r := regexp.MustCompile("[^\\s]+")
	allPairs :=r.FindAllString(text, -1)
	mapPairs:=make(map[int]int)
	for i:= 0; i<n;i++{
		intPair,_:=strconv.Atoi(allPairs[i])
		mapPairs[intPair] = mapPairs[intPair] + 1
	}
	var pairs int = 0
	for _, element := range mapPairs {
		pairs = pairs + element / 2
	}
	fmt.Println(pairs)
}