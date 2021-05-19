package main

import "fmt"
type node struct {

	A int
	Nleft *node
	Nright *node
}

type linkedList struct{

	Max int
	Root *node;
}


func insert (ll linkedList, value int) (linkedList) {

	if(ll.Root == nil) {
		ll.Root = &node{A: value, N: nil}
	}else{
		index := ll.Root;
		for index.N != nil {
			index = index.N
		}
		index.N = &node{A: value, N: nil}
	}
	ll.Max++
	return ll
}

func remove (ll linkedList) (linkedList) {


	if(ll.Root== nil){
		fmt.Println("Nada a se remover")
		return ll
	}else {
		index := ll.Root;
		if(index.N == nil){
			ll.Root = nil
			return ll
		}
		for index.N.N != nil {
			index = index.N
		}
		index.N = nil
		return ll
	}
}

func printll(ll linkedList){
	if(ll.Root == nil) {
		fmt.Println("Lista vazia")
	}else{
		index := *ll.Root;
		fmt.Print(index.A)
		for index.N != nil{
			index = *index.N
			fmt.Print(index.A)
		}
		fmt.Println(" FIM")
		
	}
	ll.Max++
	
}

func main() {

	ll :=linkedList{Max:0, Root: nil}
	ll = insert(ll, 1)
	ll = insert(ll, 2)
	ll = insert(ll, 3)
	ll = insert(ll, 4)
	printll(ll)
	ll = remove(ll)
	printll(ll)
	ll = remove(ll)
	printll(ll)
	ll = remove(ll)
	printll(ll)
	ll = remove(ll)
	printll(ll)
	ll = remove(ll)
	printll(ll)
}