package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
  )
type node struct {

	A int
	Nleft *node
	Nright *node
}

type linkedList struct{

	Max int
	Root *node;
}

func recursiveIterate (index *node, height int) (int) {

	if(index == nil){return height}
	value1:=recursiveIterate(index.Nleft, height+1)
	value2:=recursiveIterate(index.Nright, height+1)
	if(value1>value2){return value1}
	return value2
}

func getHeight (ll linkedList) (int) {

	return recursiveIterate(ll.Root, 0)
}
func insert (ll linkedList, value int) (linkedList) {
	fmt.Println("INSERINDO VALOR: ", value)
	if(ll.Root == nil) {
		ll.Root = &node{A: value, Nright: nil, Nleft: nil}
		return ll
	}else{
		index := ll.Root;
		flag  := true;
		for flag {
			if(value == index.A) {return ll}
			if(value > index.A) {
				fmt.Println("Direita")
				if(index.Nright == nil){
					index.Nright = &node{A: value, Nright: nil, Nleft: nil}
					flag = false
				}else{
					index = index.Nright;
				}
			}else {
				fmt.Println("exqierda")
				if(index.Nleft == nil){
					index.Nleft = &node{A: value, Nright: nil, Nleft: nil}
					flag = false
				}else {
					index = index.Nleft;
				}
			}
		}
	}
	ll.Max++
	return ll
}
func getInput() ([]int, int) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	n, _ := strconv.Atoi(text)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	r := regexp.MustCompile("[^\\s]+")
	arrayString :=r.FindAllString(text, -1)
	var arrayInt []int
	for _, s := range arrayString {
		value, _ :=strconv.Atoi(s)
		arrayInt = append(arrayInt, value)
	}
	fmt.Println(arrayInt)
	return arrayInt, n
}



func main() {

	ll :=linkedList{Max:0, Root: nil}
	nodes, _ := getInput()
	for _, valueNode := range nodes {
		ll = insert(ll, valueNode)
	}
	fmt.Println(getHeight(ll)-1)
}