package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'contacts' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts 2D_STRING_ARRAY queries as parameter.
 */



type trieNode struct {

	value string
	listOfLetters map[rune]*trieNode
	finalWord bool
}

type trieTree struct {
	root map[rune]*trieNode
	finalWord bool
}


func addFunc(tree *trieTree, contact string ){

	var index *trieNode;
	flagFinalWord:=false
	for i, char := range contact {
		if(i== len(contact)-1){flagFinalWord = true}
		if(i == 0){
			if(tree.root == nil){
				tree.root = make(map[rune]*trieNode)
				tree.finalWord = flagFinalWord
				tree.root[char] = &trieNode{value: contact[:i+1], listOfLetters: make(map[rune]*trieNode), finalWord: flagFinalWord}
				index = tree.root[char]
				
			}else {
				index = tree.root[char]
			}
		}else{ 
			if(index.listOfLetters[char] == nil){
				index.listOfLetters = make(map[rune]*trieNode)
				index.listOfLetters[char] = &trieNode{value: contact[:i+1], listOfLetters: nil, finalWord: flagFinalWord}
				index = index.listOfLetters[char]
			}else {
				index = index.listOfLetters[char]
			}
		}
	}
}
func searchForWords(node *trieNode, returnValue int32) int32 {
	if (node.listOfLetters == nil){
		if(node.finalWord){
			fmt.Println("entrei", node)
			returnValue++
		}
		return returnValue
	}
	for _, object := range node.listOfLetters{

		if(object.finalWord == true && object.listOfLetters != nil) {
			fmt.Println("conta1", object)
			returnValue = returnValue + 1;
		}
		returnValue = searchForWords(object, returnValue)
	}
	return returnValue

}

func findFunc(tree *trieTree, contact string) int32{
	fmt.Println("COMECOU", tree)
	fmt.Println("TEXTO", contact)
	var index *trieNode;
	returnValue:=int32(0)
	for i, char := range contact {
		if(i == 0){
			if(tree.root == nil) {
				return returnValue
			}
			
			if(tree.root[char] == nil){
				return returnValue;
			}
			index = tree.root[char];
		}else {
			if(index.listOfLetters[char] == nil ){
				return returnValue
			}
			fmt.Println("OLHA O INDICE", index)
			if(index.value == contact && index.finalWord){
				fmt.Println("AQUI")
				returnValue++

			}
			index = index.listOfLetters[char];
		}
	}
	if(index.finalWord){returnValue++}
	if(index.listOfLetters == nil) {return returnValue}
	return searchForWords(index, returnValue)
}

func contacts(queries [][]string) []int32 {
	tree := &trieTree{finalWord: false}
	var stackInt[]int32
	for _, s := range queries {
		if(s[0] == "add"){
			addFunc(tree, s[1])
		}
		if(s[0] == "find"){
			stackInt = append(stackInt, findFunc(tree,s[1]))
		}
	}
	return stackInt
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    queriesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var queries [][]string
    for i := 0; i < int(queriesRows); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []string
        for _, queriesRowItem := range queriesRowTemp {
            queriesRow = append(queriesRow, queriesRowItem)
        }

        if len(queriesRow) != 2 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := contacts(queries)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
