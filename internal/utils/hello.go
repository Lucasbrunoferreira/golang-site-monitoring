package utils

import (
	"fmt"
	"os"
)

func ShowHello() {
	name := os.Getenv("USER")
	version := 0.1
	fmt.Println("Hello", name)
	fmt.Println("This software as version", version)
}