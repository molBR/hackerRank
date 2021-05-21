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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	
	var stack []string 
	fmt.Println(s)
	for _, char := range s {
		fmt.Println(string(char))
		if(char == '{'){
			stack = append(stack, "{")

		}
		if(char == '}'){
			if(len(stack)==0) {return "NO"}
			n := len(stack) -1
			if(stack[n]== "{"){
				stack = stack[:n]
			}else{return "NO"}
			
		}

		if(char == '['){
			stack = append(stack, "[")
		}
		if(char == ']' ){
			if(len(stack)==0) {return "NO"}
			n := len(stack)-1
			if(stack[n] == "[") {
				stack = stack[:n]
			}else{return "NO"}
			
		}

		if(char == '('){
			stack = append(stack, "(")
		}
		if(char == ')'){
			if(len(stack)==0) {return "NO"}
			n := len(stack) -1
			if(stack[n]=="("){
				stack = stack[:n]
			}else{return "NO"}
		}
	}
	if(len(stack)==0){
		return "YES"
	}
	return "NO"
	

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        s := readLine(reader)

        result := isBalanced(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
