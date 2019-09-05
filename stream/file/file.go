package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	fileUri = "d:/go_project/src/github.com/max_workspace/golang_learn/stream/input/input.dat"
)

func TestReadFromFile() {
	fmt.Println("test read from file")
	inputFile, inputError := os.Open(fileUri)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		dir, ok := os.Getwd()
		fmt.Printf("dir : %v\n status: %v\n", dir, ok)
		fmt.Printf("description: %v\n", inputError)
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
	fmt.Println("\n")
}
