package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "I am josheph merino"
	file, err := os.Create("./newFile")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Written in file total characters ", length)
	defer file.Close()
	readFile("./newFile")
}

func readFile(inputFile string) {
	bytesArr, err := os.ReadFile(inputFile)
	checkError(err)
	fmt.Println("data of file is ", string(bytesArr))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
