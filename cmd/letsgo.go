package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yousafgill/letsgo/pkg/controlflow"
	"github.com/yousafgill/letsgo/pkg/encryption"
	"github.com/yousafgill/letsgo/pkg/function"
	"github.com/yousafgill/letsgo/pkg/goroutine"
	"github.com/yousafgill/letsgo/pkg/loop"
	"github.com/yousafgill/letsgo/pkg/operator"
	"github.com/yousafgill/letsgo/pkg/variable"
)

func main() {

	//clear console
	print("\033[H\033[2J")

	//1: Please choose an option
	println(`Please choose an option:
1: Print Hello World
2: Print Operators 
3: Print Variables
4: Print Control Flow
5: Print Functions 
6: Print Loops
7: Test Go Routines
8: Test Encryption
99: Exit`)

	//2: Read user input
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//3: Check user input
	switch strings.ToLower(input) {
	case "1":
		println("Hello World")
	case "2":
		operator.PrintOperators()
	case "3":
		variable.PrintVariables()
	case "4":
		controlflow.PrintControlFlow()
	case "5":
		function.PrintFunctions()
	case "6":
		loop.PrintLoops()
	case "7":
		goroutine.TestGoRoutine()
	case "8":
		encryption.RunEncryption()
	case "99":
		return
	default:
		println("Invalid input")
	}

	//write a function which count to 1 billion and print the time it took

	

}