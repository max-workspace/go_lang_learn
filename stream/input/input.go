package input

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName string
	age                 int
	inputString         = "56.12 0+ 5212 @ Go"
	inputFormat         = "%f 0+ %d @ %s"
	f                   float32
	i                   int
	s                   string
	inputReader         *bufio.Reader
	err                 error
	input               string
)

func TestInputFromUser() {
	fmt.Println("test input from user")
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Println("Please enter your name and age:(like join 20)")
	fmt.Scanf("%s %d", &firstName, &age)
	fmt.Printf("Hi %s, your age is %02d!\n", firstName, age)
	fmt.Println("\n")
}

func TestInputFromString() {
	fmt.Println("test input from string")

	fmt.Sscanln(inputString, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)

	fmt.Sscanf(inputString, inputFormat, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)

	fmt.Println("\n")
}

func TestInputFromBuffer() {
	fmt.Println("test input from buffer")

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
	fmt.Println("\n")
}
