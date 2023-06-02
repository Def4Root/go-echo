package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("")
		return
	}

	args := []string{}
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	for _, arg := range args {
		fmt.Print(arg, " ")
	}

	fmt.Println("")
}