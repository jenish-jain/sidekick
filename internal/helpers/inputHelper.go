package helpers

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

type Input interface {
	GetLineInput(inputQuery string) []byte
	GetPasswordInput(inputQuery string) []byte
}

func GetLineInput(inputQuery string) []byte {
	reader := bufio.NewReader(os.Stdin)
	color.Blue(inputQuery)
	userInput, _, error := reader.ReadLine()
	if error != nil {
		fmt.Errorf("error reading user input for input query %s , error : %s", inputQuery, error.Error())
		panic(error)
	}
	return userInput
}

func GetPasswordInput(inputQuery string) []byte {
	color.Blue(inputQuery)
	bytePassword, error := terminal.ReadPassword(syscall.Stdin)
	if error != nil {
		fmt.Errorf("error reading passowrd input for input query %s , error : %s", inputQuery, error.Error())
		panic(error)
	}
	return bytePassword
}
