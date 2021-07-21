package utils

import (
	"fmt"
	"os"
)

func ShowErrorMessage() {
	fmt.Println("Please, select a valid option: 1, 2 or 0")
	os.Exit(-1)
}